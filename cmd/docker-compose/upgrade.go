package main

import (
	dockercompose "get.porter.sh/mixin/docker-compose/pkg/docker-compose"
	"github.com/spf13/cobra"
)

func buildUpgradeCommand(m *dockercompose.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute(cmd.Context())
		},
	}
	return cmd
}
