package dockercompose

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	const buildOutput = `RUN apt-get update && apt-get install -y curl && \
curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose && \
chmod +x /usr/local/bin/docker-compose
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
