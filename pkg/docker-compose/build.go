package dockercompose

import (
	"fmt"
)

const dockerComposeVersion = "1.25.4"

// Build installs the docker and docker-compose binaries
func (m *Mixin) Build() error {
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
