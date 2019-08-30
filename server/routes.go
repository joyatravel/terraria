package server

import (
	info "github.com/tugolo/terraria/internal/info"
	srvinfo "github.com/tugolo/terraria/server/internal/info"
	"github.com/tugolo/terraria/terraria/http"

	"github.com/tugolo/terraria/pkg/httputil"
)

func (srv *Server) registerRoutes() {
	e := srv.echo
	svcs := &srv.services

	// Register error and info handlers.
	e.HTTPErrorHandler = httputil.ErrorHandler(
		srv.log.WithField("handler", "error"),
	)
	e.GET("/", httputil.InfoHandler(srvinfo.ServerName, info.Version))

	// Register Terraria handlers.
	e.GET("/world", http.WorldHandler(svcs.World, svcs.Status))
	e.GET("/status", http.StatusHandler(svcs.Status))
}
