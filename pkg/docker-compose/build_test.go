package dockercompose

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	const buildOutput = `ADD --chmod=755 https://github.com/docker/compose/releases/download/v2.10.2/docker-compose-linux-x86_64 /usr/local/lib/docker/cli-plugins/docker-compose
ENV PATH="$PATH:/usr/local/lib/docker/cli-plugins"
`

	t.Run("build", func(t *testing.T) {
		m := NewTestMixin(t)
		m.DebugMode = false

		err := m.Build(context.Background())
		require.NoError(t, err, "build failed")

		wantOutput := buildOutput
		gotOutput := m.TestContext.GetOutput()
		assert.Equal(t, wantOutput, gotOutput)
	})

	t.Run("build with a defined docker-compose client version", func(t *testing.T) {
		b, err := ioutil.ReadFile("testdata/build-input-with-version.yaml")
		require.NoError(t, err)

		m := NewTestMixin(t)
		m.DebugMode = false
		m.In = bytes.NewReader(b)

		err = m.Build(context.Background())
		require.NoError(t, err, "build failed")

		wantOutput := buildOutput
		gotOutput := m.TestContext.GetOutput()
		assert.Equal(t, wantOutput, gotOutput)
	})

	t.Run("build with invalid config", func(t *testing.T) {
		b, err := ioutil.ReadFile("testdata/build-input-with-invalid-config.yaml")
		require.NoError(t, err)

		m := NewTestMixin(t)
		m.DebugMode = false
		m.In = bytes.NewReader(b)

		err = m.Build(context.Background())
		require.Error(t, err, "build failed")
	})
}
