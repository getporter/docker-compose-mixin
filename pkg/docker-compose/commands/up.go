package commands

import "get.porter.sh/porter/pkg/exec/builder"

var _ Command = UpCommand{}

type UpCommand struct {
	Arguments      []string      `yaml:"arguments,omitempty"`
	Flags          builder.Flags `yaml:"flags,omitempty"`
	Outputs        []Output      `yaml:"outputs,omitempty"`
	SuppressOutput bool          `yaml:"suppress-output,omitempty"`
}

func (c UpCommand) GetCommand() string {
	return "up"
}

func (c UpCommand) GetArguments() []string {
	// Always use suffix arguments
	return nil
}

func (c UpCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c UpCommand) GetSuffixArguments() []string {
	return c.Arguments
}

func (c UpCommand) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(c.Outputs))
	for i := range c.Outputs {
		outputs[i] = c.Outputs[i]
	}
	return outputs
}

func (c UpCommand) SuppressesOutput() bool {
	return c.SuppressOutput
}

func (c UpCommand) GetWorkingDir() string {
	return "."
}
