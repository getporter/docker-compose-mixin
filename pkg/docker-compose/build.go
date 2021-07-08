package dockercompose

import (
	"fmt"

	"get.porter.sh/porter/pkg/exec/builder"
	yaml "gopkg.in/yaml.v2"
)

const dockerComposeDefaultVersion = "1.29.2"

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
func (m *Mixin) Build() error {
	// Create new Builder.
	var input BuildInput
	err := builder.LoadAction(m.Context, "", func(contents []byte) (interface{}, error) {
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

	dockerfileLines := fmt.Sprintf(`ENV DOCKER_VERSION="19.03.8"
RUN apt-get update && apt-get install -y python3-pip wget && pip3 install --upgrade pip && \
  wget https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz && \
  tar -xvf docker-${DOCKER_VERSION}.tgz && \
  mv docker/docker /usr/bin/docker && \
  chmod +x /usr/bin/docker && \
  rm -rf docker/ docker-${DOCKER_VERSION}.tgz && \
  pip3 install docker-compose==%s`, dockerComposeVersion)

	fmt.Fprintln(m.Out, dockerfileLines)

	return nil
}
