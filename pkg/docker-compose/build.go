package dockercompose

import (
	"context"
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	"gopkg.in/yaml.v3"
)

const dockerComposeDefaultVersion = "2.10.2"

// BuildInput represents stdin passed to the mixin for the build command.
type BuildInput struct {
	Config MixinConfig
}

// MixinConfig represents configuration that can be set on the docker-compose mixin in porter.yaml
// mixins:
// - docker-compose:
//     clientVersion: 1.29.2

type MixinConfig struct {
	ClientVersion string `yaml:"clientVersion,omitempty"`
}

// Build installs the docker and docker-compose binaries
func (m *Mixin) Build(ctx context.Context) error {
	// Create new Builder.
	var input BuildInput
	err := builder.LoadAction(ctx, m.RuntimeConfig, "", func(contents []byte) (interface{}, error) {
		err := yaml.Unmarshal(contents, &input)
		return &input, err
	})
	if err != nil {
		return err
	}

	var dockerComposeVersion string
	if input.Config.ClientVersion != "" {
		dockerComposeVersion = input.Config.ClientVersion
	} else {
		dockerComposeVersion = dockerComposeDefaultVersion
	}

	dockerfileLines := fmt.Sprintf(`ADD --chmod=755 https://github.com/docker/compose/releases/download/v%s/docker-compose-linux-x86_64 /usr/local/lib/docker/cli-plugins/docker-compose
ENV PATH="$PATH:/usr/local/lib/docker/cli-plugins"`, dockerComposeVersion)

	fmt.Fprintln(m.Out, dockerfileLines)

	return nil
}
