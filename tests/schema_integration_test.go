package tests

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

// Add a test that checked that the schema was packed into the binary
// properly. Requires a make clean xbuild-all first.
func TestSchema(t *testing.T) {
	schemaBackup := "../pkg/docker-compose/schema/schema.json.bak"
	schemaPath := "../pkg/docker-compose/schema/schema.json"
	defer os.Rename(schemaBackup, schemaPath)
	err := os.Rename(schemaPath, schemaBackup)
	require.NoError(t, err, "failed to sabotage the schema.json")

	output := &bytes.Buffer{}
	cmd := exec.Command("docker-compose", "schema")
	cmd.Path = "../bin/mixins/docker-compose/docker-compose"
	cmd.Stdout = output
	cmd.Stderr = output

	err = cmd.Start()
	require.NoError(t, err, "failed to start the docker-compose schema command")

	err = cmd.Wait()
	t.Log(output)
	require.NoError(t, err, "docker-compose schema failed")
}
