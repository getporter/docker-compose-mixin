package main

import (
	dockercompose "get.porter.sh/mixin/docker-compose/pkg/docker-compose"
	"github.com/spf13/cobra"
)

func buildSchemaCommand(m *dockercompose.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema",
		Short: "Print the json schema for the mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.PrintSchema()
		},
	}
	return cmd
}
