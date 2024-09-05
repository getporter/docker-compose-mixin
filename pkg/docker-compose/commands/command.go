package commands

import "get.porter.sh/porter/pkg/exec/builder"

type Command interface {
	builder.ExecutableStep
	builder.HasOrderedArguments
	builder.StepWithOutputs
	builder.SuppressesOutput
}

var _ builder.OutputJsonPath = Output{}
var _ builder.OutputFile = Output{}

type Output struct {
	Name     string `yaml:"name"`
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
