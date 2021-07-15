package dockercompose

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	const buildOutput = `RUN apt-get update && apt-get install -y curl && \
curl -L "https://github.com/docker/compose/releases/download/1.26.0/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose && \
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
}
