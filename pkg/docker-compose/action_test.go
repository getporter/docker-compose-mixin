package dockercompose

import (
	"io/ioutil"
	"testing"

	"get.porter.sh/porter/pkg/exec/builder"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	yaml "gopkg.in/yaml.v2"
)

func TestMixin_UnmarshalStep(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, "Compose Up", step.Description)
	assert.NotEmpty(t, step.Outputs)
	assert.Equal(t, Output{Name: "containerId", JsonPath: "$Id"}, step.Outputs[0])

	require.Len(t, step.Arguments, 2)
	assert.Equal(t, "up", step.Arguments[0])
	assert.Equal(t, "-d", step.Arguments[1])

	require.Len(t, step.Flags, 1)
	assert.Equal(t, builder.NewFlag("timeout", "25"), step.Flags[0])

	assert.Equal(t, false, step.SuppressOutput)
	assert.Equal(t, false, step.SuppressesOutput())
}

func TestStep_SuppressesOutput(t *testing.T) {
	b, err := ioutil.ReadFile("testdata/step-input-suppress-output.yaml")
	require.NoError(t, err)

	var action Action
	err = yaml.Unmarshal(b, &action)
	require.NoError(t, err)
	require.Len(t, action.Steps, 1)

	step := action.Steps[0]
	assert.Equal(t, true, step.SuppressOutput)
	assert.Equal(t, true, step.SuppressesOutput())
}
