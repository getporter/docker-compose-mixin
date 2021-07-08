package dockercompose

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	const buildOutput = `ENV DOCKER_VERSION="19.03.8"
RUN apt-get update && apt-get install -y python3-pip wget && pip3 install --upgrade pip && \
  wget https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz && \
  tar -xvf docker-${DOCKER_VERSION}.tgz && \
  mv docker/docker /usr/bin/docker && \
  chmod +x /usr/bin/docker && \
  rm -rf docker/ docker-${DOCKER_VERSION}.tgz && \
  pip3 install docker-compose==1.29.2
`

	t.Run("build", func(t *testing.T) {
		m := NewTestMixin(t)
		m.Debug = false

		err := m.Build()
		require.NoError(t, err, "build failed")

		wantOutput := buildOutput
		gotOutput := m.TestContext.GetOutput()
		assert.Equal(t, wantOutput, gotOutput)
	})

	t.Run("build with a defined docker-compose client version", func(t *testing.T) {
		b, err := ioutil.ReadFile("testdata/build-input-with-version.yaml")
		require.NoError(t, err)

		m := NewTestMixin(t)
		m.Debug = false
		m.In = bytes.NewReader(b)

		err = m.Build()
		require.NoError(t, err, "build failed")

		wantOutput := buildOutput
		gotOutput := m.TestContext.GetOutput()
		assert.Equal(t, wantOutput, gotOutput)
	})

	t.Run("build with invalid config", func(t *testing.T) {
		b, err := ioutil.ReadFile("testdata/build-input-with-invalid-config.yaml")
		require.NoError(t, err)

		m := NewTestMixin(t)
		m.Debug = false
		m.In = bytes.NewReader(b)

		err = m.Build()
		require.Error(t, err, "build failed")
	})
}
