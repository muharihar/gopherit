package cmd

import (
	"00-newapp-template/internal/app/cmd/client"
	"00-newapp-template/internal/pkg"
	"00-newapp-template/internal/pkg/adapter"
	"00-newapp-template/internal/pkg/ui"
	"fmt"
	"github.com/spf13/cobra"
)

type Client struct {
	Config  *pkg.Config
	Adapter *adapter.Adapter
	CLI     ui.CLI
}

func NewClient(config *pkg.Config) (c Client) {
	c.Config = config
	c.CLI = ui.NewCLI(config)
	c.Adapter = adapter.NewAdapter(c.Config)
	return
}

func (c *Client) Client(cmd *cobra.Command, args []string) {
	cmd.Help()
	return
}
func (c *Client) List(cmd *cobra.Command, args []string) {
	client.List(c.Adapter, c.CLI)
	return
}

func (c *Client) Delete(cmd *cobra.Command, args []string) {
	client.Delete(c.Adapter, c.CLI)
	return
}

func (c *Client) Update(cmd *cobra.Command, args []string) {
	fmt.Printf("UpdateCommand\n")
	return
}
