package dockercompose

import (
	"fmt"
)

const dockerComposeVersion = "1.26.0"

// Build installs the docker and docker-compose binaries
func (m *Mixin) Build() error {
	dockerfileLines := fmt.Sprintf(`RUN apt-get update && apt-get install -y curl && \
curl -L "https://github.com/docker/compose/releases/download/%s/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose && \
chmod +x /usr/local/bin/docker-compose`, dockerComposeVersion)

	fmt.Fprintln(m.Out, dockerfileLines)

	return nil
}
