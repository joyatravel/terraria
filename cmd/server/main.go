package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/tugolo/terraria/terraria/gameserver"

	"github.com/tugolo/terraria/terraria"

	"github.com/tugolo/terraria/terraria/disk"

	"go.stevenxie.me/gopkg/cmdutil"
	"go.stevenxie.me/gopkg/configutil"
	"go.stevenxie.me/guillotine"

	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	cmdinfo "github.com/tugolo/terraria/cmd/server/internal/info"
	info "github.com/tugolo/terraria/internal/info"

	"github.com/tugolo/terraria/cmd/server/config"
	cmdutilx "github.com/tugolo/terraria/pkg/cmdutil"
	"github.com/tugolo/terraria/server"
)

func main() {
	// Load envvars from dotenv.
	if err := configutil.LoadEnv(); err != nil {
		cmdutil.Fatalf("Failed to load dotenv file: %v\n", err)
	}

	app := cli.NewApp()
	app.Name = cmdinfo.Name
	app.Usage = "The public API server that fronts the Tugolo Terraria server."
	app.UsageText = "server [global options]"
	app.Version = info.Version
	app.Action = run

	// Hide help command.
	app.Commands = []cli.Command{{Name: "help", Hidden: true}}

	// Configure flags.
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "port",
			Usage:       "port that the server listens on",
			Value:       3000,
			Destination: &flags.Port,
		},
		cli.BoolFlag{
			Name:  "help,h",
			Usage: "show help",
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v\n", err)
		os.Exit(1)
	}
}

var flags struct {
	Port int
}

func run(*cli.Context) (err error) {
	// Init logger, load config.
	log := cmdutilx.Logger()
	cfg, err := config.Load()
	if err != nil {
		return errors.Wrap(err, "loading config")
	}

	// Init guillotine.
	guillo := guillotine.New(
		guillotine.WithLogger(log.WithField("component", "guillotine.Guillotine")),
	)
	guillo.TriggerOnTerminate()
	defer func() {
		if ok, _ := guillo.Execute(); !ok && (err != nil) {
			err = errors.New("guillotine finished running with errors")
		}
	}()

	// Connect to data sources.
	log.Info("Connecting to data sources...")

	// Connect to gameserver.
	gsclient := gameserver.NewClient(cfg.Terraria.Address)
	if err = gsclient.Ping(); err != nil {
		return errors.Wrap(err, "pinging gameserver")
	}

	// Init services.
	log.Info("Initializing services...")

	var worlds terraria.WorldService
	worlds = disk.NewWorldFinder(
		cfg.Terraria.WorldDir,
		func(cfg *disk.WorldFinderConfig) {
			cfg.Logger = log.WithField("component", "disk.WorldFinder")
		},
	)

	var status terraria.StatusService
	status = gameserver.NewStatusService(
		gsclient,
		func(cfg *gameserver.StatusServiceConfig) {
			cfg.Logger = log.WithField("component", "gameserver.StatusService")
		},
	)

	// Start HTTP server.
	log.Info("Initializing HTTP server...")
	srv := server.New(
		server.Services{
			World:  worlds,
			Status: status,
		},
		server.WithLogger(log.WithField("component", "server.Server")),
	)

	guillo.AddFinalizer(func() error {
		var (
			entry = logrus.NewEntry(log)
			ctx   = context.Background()
		)
		if timeout := cfg.Server.ShutdownTimeout; timeout != nil {
			entry = entry.WithField("timeout", *timeout)
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, *timeout)
			defer cancel()
		}
		entry.Info("Shutting down HTTP server...")
		err := srv.Shutdown(ctx)
		return errors.Wrap(err, "shutting down server")
	})

	err = srv.ListenAndServe(fmt.Sprintf(":%d", flags.Port))
	if !errors.Is(err, http.ErrServerClosed) {
		guillo.Trigger()
		return errors.Wrap(err, "starting HTTP server")
	}

	return nil
}
