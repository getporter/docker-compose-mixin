package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	dockercompose "get.porter.sh/mixin/docker-compose/pkg/docker-compose"
	"get.porter.sh/porter/pkg/cli"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	run := func() int {
		ctx := context.Background()
		m := dockercompose.New()
		ctx, err := m.ConfigureLogging(ctx)
		if err != nil {
			fmt.Println(err)
			os.Exit(cli.ExitCodeErr)
		}
		cmd := buildRootCommand(m, os.Stdin)

		// We don't have tracing working inside a bundle working currently.
		// We are using StartRootSpan anyway because it creates a TraceLogger and sets it
		// on the context, so we can grab it later
		ctx, log := m.StartRootSpan(ctx, "docker-compose")
		defer func() {
			// Capture panics and trace them
			if panicErr := recover(); panicErr != nil {
				log.Error(fmt.Errorf("%s", panicErr),
					attribute.Bool("panic", true),
					attribute.String("stackTrace", string(debug.Stack())))
				log.EndSpan()
				m.Close()
				os.Exit(cli.ExitCodeErr)
			} else {
				log.Close()
				m.Close()
			}
		}()

		if err := cmd.ExecuteContext(ctx); err != nil {
			return cli.ExitCodeErr
		}
		return cli.ExitCodeSuccess
	}
	os.Exit(run())
}

func buildRootCommand(m *dockercompose.Mixin, in io.Reader) *cobra.Command {
	cmd := &cobra.Command{
		Use:  "docker-compose",
		Long: "Use the Docker Compose CLI with Porter",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Enable swapping out stdout/stderr for testing
			m.In = in
			m.Out = cmd.OutOrStdout()
			m.Err = cmd.OutOrStderr()
		},
		SilenceUsage: true,
	}

	cmd.PersistentFlags().BoolVar(&m.DebugMode, "debug", false, "Enable debug logging")

	cmd.AddCommand(buildVersionCommand(m))
	cmd.AddCommand(buildSchemaCommand(m))
	cmd.AddCommand(buildBuildCommand(m))
	cmd.AddCommand(buildInstallCommand(m))
	cmd.AddCommand(buildInvokeCommand(m))
	cmd.AddCommand(buildUpgradeCommand(m))
	cmd.AddCommand(buildUninstallCommand(m))

	return cmd
}
