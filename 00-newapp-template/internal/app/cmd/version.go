package cmd

import (
	"00-newapp-template/internal/pkg"
	"00-newapp-template/internal/pkg/ui"
	"github.com/spf13/cobra"
)

type Version struct {
	Config *pkg.Config
	CLI    *ui.CLI
}

func NewVersion(c *pkg.Config) (v Version) {
	v.Config = c
	return
}

func (v *Version) Version(cmd *cobra.Command, args []string) {

	cli := ui.NewCLI(v.Config)

	cli.DrawGopher()

	return
}
