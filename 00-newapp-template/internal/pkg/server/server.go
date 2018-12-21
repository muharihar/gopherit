package server

import (
	"context"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	Context  context.Context
	Router   chi.Router
	HTTP     *http.Server
	Log      *log.Logger
	Finished context.CancelFunc
	DB       SimpleDB
}

func NewServer(context context.Context, listenPort string, log *log.Logger) (server Server) {
	server.Context = context
	server.Router = chi.NewRouter()
	server.HTTP = &http.Server{Addr: ":" + listenPort, Handler: server.Router}
	server.Log = log
	server.DB = NewDB()
	return
}

func (server *Server) ListenAndServe() (err error) {
	server.hookShutdownSignal()

	go func() {
		server.Log.Infof("server starting")
		err = server.HTTP.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			server.Log.Errorf("error serving: %+v", err)
		}
		server.Finished()
	}()

	select {
	case <-server.Context.Done():
		server.Log.Infof("server stopped")
	}

	return
}
func (server *Server) hookShutdownSignal() {
	stop := make(chan os.Signal)

	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	server.Context, server.Finished = context.WithCancel(server.Context)
	go func() {
		sig := <-stop
		server.Log.Infof("server terminatation signal '%s' received", sig)
		server.Finished()
	}()

	return
}
