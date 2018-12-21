package server

import (
	"00-newapp-template/internal/pkg/server"
	"context"
	log "github.com/sirupsen/logrus"
)

func Start(context context.Context, port string, log *log.Logger) {
	s := server.NewServer(context, port, log)
	s.NewRouter()
	s.ListenAndServe()

	return
}
