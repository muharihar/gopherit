package cmd

import (
	"00-newapp-template/internal/app/cmd/server"
	"00-newapp-template/internal/pkg"
	"00-newapp-template/internal/pkg/ui"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Server struct {
	Config *pkg.Config
	CLI    ui.CLI
}

func NewServer(config *pkg.Config) (c Server) {
	c.Config = config
	c.CLI = ui.NewCLI(config)
	return
}

func (c *Server) Server(cmd *cobra.Command, args []string) {
	cmd.Help()
	return
}
func (c *Server) Start(cmd *cobra.Command, args []string) {

	log := c.Config.Log.WithFields(log.Fields{
		"docroot": c.Config.Server.RootFolder,
		"port":    c.Config.Server.ListenPort,
	})

	log.Info("starting server")

	server.Start(c.Config.Context, c.Config.Server.ListenPort, c.Config.Log)

	log.Info("server finished")

	return
}
func (c *Server) Stop(cmd *cobra.Command, args []string) {
	fmt.Printf("Stop Command\n")
	return
}
