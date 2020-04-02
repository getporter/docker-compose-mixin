package dockercompose

import (
	"get.porter.sh/porter/pkg/exec/builder"
)

var _ builder.ExecutableAction = Action{}
var _ builder.BuildableAction = Action{}

type Action struct {
	Name  string
	Steps []Step // using UnmarshalYAML so that we don't need a custom type per action
}

// MakeSteps builds a slice of Steps for data to be unmarshaled into.
func (a Action) MakeSteps() interface{} {
	return &[]Step{}
}

// UnmarshalYAML takes any yaml in this form
// ACTION:
// - docker-compose: ...
// and puts the steps into the Action.Steps field
func (a *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	results, err := builder.UnmarshalAction(unmarshal, a)
	if err != nil {
		return err
	}

	for actionName, action := range results {
		a.Name = actionName
		for _, result := range action {
			step := result.(*[]Step)
			a.Steps = append(a.Steps, *step...)
		}
		break // There is only 1 action
	}
	return nil
}

func (a Action) GetSteps() []builder.ExecutableStep {
	// Go doesn't have generics, nothing to see here...
	steps := make([]builder.ExecutableStep, len(a.Steps))
	for i := range a.Steps {
		steps[i] = a.Steps[i]
	}

	return steps
}

type Step struct {
	Instruction `yaml:"docker-compose"`
}

var _ builder.ExecutableStep = Step{}
var _ builder.StepWithOutputs = Step{}
var _ builder.SuppressesOutput = Step{}

type Instruction struct {
	Name           string        `yaml:"name"`
	Description    string        `yaml:"description"`
	Arguments      []string      `yaml:"arguments,omitempty"`
	Flags          builder.Flags `yaml:"flags,omitempty"`
	Outputs        []Output      `yaml:"outputs,omitempty"`
	SuppressOutput bool          `yaml:"suppress-output,omitempty"`
}

func (s Step) GetCommand() string {
	return "docker-compose"
}

func (s Step) GetArguments() []string {
	return s.Arguments
}

func (s Step) GetFlags() builder.Flags {
	return s.Flags
}

func (s Step) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(s.Outputs))
	for i := range s.Outputs {
		outputs[i] = s.Outputs[i]
	}
	return outputs
}

func (s Step) SuppressesOutput() bool {
	return s.SuppressOutput
}

var _ builder.OutputJsonPath = Output{}
var _ builder.OutputFile = Output{}

type Output struct {
	Name string `yaml:"name"`

	// See https://porter.sh/mixins/exec/#outputs
	JsonPath string `yaml:"jsonPath,omitempty"`
	FilePath string `yaml:"path,omitempty"`
}

func (o Output) GetName() string {
	return o.Name
}

func (o Output) GetJsonPath() string {
	return o.JsonPath
}

func (o Output) GetFilePath() string {
	return o.FilePath
}
