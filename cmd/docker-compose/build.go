package main

import (
	dockercompose "get.porter.sh/mixin/docker-compose/pkg/docker-compose"
	"github.com/spf13/cobra"
)

func buildBuildCommand(m *dockercompose.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Generate Dockerfile lines for the bundle invocation image",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Build(cmd.Context())
		},
	}
	return cmd
}
