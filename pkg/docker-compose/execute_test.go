package dockercompose

import (
	"bytes"
	"context"
	"io/ioutil"
	"path"
	"testing"

	"get.porter.sh/porter/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	test.TestMainWithMockedCommandHandlers(m)
}

func TestMixin_Execute(t *testing.T) {
	testcases := []struct {
		name        string
		file        string
		wantOutput  string
		wantCommand string
	}{
		{
			"install", "testdata/install-input.yaml", "",
			"docker-compose up --build --scale 2",
		},
		{
			"down", "testdata/commands/down-input.yaml", "containerId",
			"docker-compose --file test.yml down --remove-orphans --timeout 25 serviceA serviceB",
		},
		{
			"pull", "testdata/commands/pull-input.yaml", "containerId",
			"docker-compose --file test.yml pull --ignore-pull-failures --policy missing serviceA serviceB",
		},
		{
			"up", "testdata/commands/up-input.yaml", "containerId",
			"docker-compose --file test.yml up --detach --timeout 25 serviceA serviceB",
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			m := NewTestMixin(t)

			m.Setenv(test.ExpectedCommandEnv, tc.wantCommand)
			mixinInputB, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			m.In = bytes.NewBuffer(mixinInputB)

			err = m.Execute(context.Background())
			require.NoError(t, err, "execute failed")

			if tc.wantOutput == "" {
				outputs, _ := m.FileSystem.ReadDir("/cnab/app/porter/outputs")
				assert.Empty(t, outputs, "expected no outputs to be created")
			} else {
				wantPath := path.Join("/cnab/app/porter/outputs", tc.wantOutput)
				exists, _ := m.FileSystem.Exists(wantPath)
				assert.True(t, exists, "output file was not created %s", wantPath)
			}
		})
	}
}
