package commands

import "get.porter.sh/porter/pkg/exec/builder"

var _ Command = DownCommand{}

type DownCommand struct {
	Arguments      []string      `yaml:"arguments,omitempty"`
	Flags          builder.Flags `yaml:"flags,omitempty"`
	Outputs        []Output      `yaml:"outputs,omitempty"`
	SuppressOutput bool          `yaml:"suppress-output,omitempty"`
}

func (c DownCommand) GetCommand() string {
	return "down"
}

func (c DownCommand) GetArguments() []string {
	// Always use suffix arguments
	return nil
}

func (c DownCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c DownCommand) GetSuffixArguments() []string {
	return c.Arguments
}

func (c DownCommand) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(c.Outputs))
	for i := range c.Outputs {
		outputs[i] = c.Outputs[i]
	}
	return outputs
}

func (c DownCommand) SuppressesOutput() bool {
	return c.SuppressOutput
}

func (c DownCommand) GetWorkingDir() string {
	return "."
}
