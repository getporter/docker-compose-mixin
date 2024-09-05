package dockercompose

import (
	"io/ioutil"
	"testing"

	"get.porter.sh/mixin/docker-compose/pkg/docker-compose/commands"
	"get.porter.sh/porter/pkg/exec/builder"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v3"
)

func TestMixin_UnmarshalStep(t *testing.T) {
	testcases := []struct {
		name            string            // Test case name
		file            string            // Path to the test input yaml
		wantDescription string            // Description that you expect to be found
		wantArguments   []string          // Arguments that you expect to be found
		wantFlags       builder.Flags     // Flags that you expect to be found
		wantSuffixArgs  []string          // Suffix arguments that you expect to be found
		wantOutputs     []commands.Output // Outputs that you expect to be found
		wantSuppress    bool
	}{
		{
			"step", "testdata/step-input.yaml", "Compose Up",
			[]string{"up", "-d"},
			builder.Flags{builder.NewFlag("timeout", "25")}, nil,
			[]commands.Output{{Name: "containerId", JsonPath: "$Id"}}, false,
		},
		{
			"step-suppress-output", "testdata/step-input-suppress-output.yaml",
			"Supressed Surprise", []string{"surprise", "me"}, nil, nil, nil, true,
		},
		{
			"down", "testdata/commands/down-input.yaml", "Compose Down", nil,
			builder.Flags{builder.NewFlag("file", "test.yml")},
			[]string{"down", "--remove-orphans", "--timeout", "25", "serviceA", "serviceB"},
			[]commands.Output{{Name: "containerId", JsonPath: "$Id"}}, false,
		},
		{
			"pull", "testdata/commands/pull-input.yaml", "Compose Pull", nil,
			builder.Flags{builder.NewFlag("file", "test.yml")},
			[]string{"pull", "--ignore-pull-failures", "--policy", "missing", "serviceA", "serviceB"},
			[]commands.Output{{Name: "containerId", JsonPath: "$Id"}}, false,
		},
		{
			"up", "testdata/commands/up-input.yaml", "Compose Up", nil,
			builder.Flags{builder.NewFlag("file", "test.yml")},
			[]string{"up", "--detach", "--timeout", "25", "serviceA", "serviceB"},
			[]commands.Output{{Name: "containerId", JsonPath: "$Id"}}, false,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			b, err := ioutil.ReadFile(tc.file)
			require.NoError(t, err)

			var action Action
			err = yaml.Unmarshal(b, &action)
			require.NoError(t, err)
			require.Len(t, action.Steps, 1)

			step := action.Steps[0]
			assert.Equal(t, tc.wantDescription, step.Description)

			args := step.GetArguments()
			assert.Equal(t, tc.wantArguments, args)

			flags := step.GetFlags()
			assert.Equal(t, tc.wantFlags, flags)

			suffixArgs := step.GetSuffixArguments()
			assert.Equal(t, tc.wantSuffixArgs, suffixArgs)

			outputs := step.GetOutputs()
			assert.ElementsMatch(t, tc.wantOutputs, outputs)

			assert.Equal(t, tc.wantSuppress, step.SuppressesOutput(), "invalid suppress-output")
		})
	}
}
