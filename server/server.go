package server

import (
	"context"
	"io/ioutil"
	"os"

	"go.stevenxie.me/gopkg/zero"

	"github.com/cockroachdb/errors"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"

	"github.com/tugolo/terraria/terraria"
)

// New creates a new Server.
func New(s Services, opts ...Option) *Server {
	cfg := Config{
		Logger: zero.Logger(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	// Configure echo.
	echo := echo.New()
	echo.Logger.SetOutput(ioutil.Discard) // disable logger

	// Enable Access-Control-Allow-Origin: * during development.
	if os.Getenv("GOENV") == "development" {
		echo.Use(middleware.CORS())
	}

	// Build and configure server.
	return &Server{
		echo:     echo,
		log:      cfg.Logger,
		services: s,
	}
}

type (
	// Server serves a REST API for user account management.
	Server struct {
		echo     *echo.Echo
		log      logrus.FieldLogger
		services Services
	}

	// Config configures a Server.
	Config struct {
		Logger logrus.FieldLogger
	}

	// Services are used to handle requests.
	Services struct {
		World  terraria.WorldService
		Status terraria.StatusService
	}
)

// ListenAndServe listens and serves on the specified address.
func (srv *Server) ListenAndServe(addr string) error {
	if addr == "" {
		return errors.New("http: address must be non-empty")
	}
	srv.registerRoutes()

	// Listen for connections.
	srv.log.WithField("addr", addr).Info("Listening for connections...")
	return srv.echo.Start(addr)
}

// Shutdown shuts down the server gracefully without interupting any active
// connections.
func (srv *Server) Shutdown(ctx context.Context) error {
	return srv.echo.Shutdown(ctx)
}
