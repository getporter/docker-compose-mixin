package commands

import "get.porter.sh/porter/pkg/exec/builder"

var _ Command = PullCommand{}

type PullCommand struct {
	Arguments      []string      `yaml:"arguments,omitempty"`
	Flags          builder.Flags `yaml:"flags,omitempty"`
	Outputs        []Output      `yaml:"outputs,omitempty"`
	SuppressOutput bool          `yaml:"suppress-output,omitempty"`
}

func (c PullCommand) GetCommand() string {
	return "pull"
}

func (c PullCommand) GetArguments() []string {
	// Always use suffix arguments
	return nil
}

func (c PullCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c PullCommand) GetSuffixArguments() []string {
	return c.Arguments
}

func (c PullCommand) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(c.Outputs))
	for i := range c.Outputs {
		outputs[i] = c.Outputs[i]
	}
	return outputs
}

func (c PullCommand) SuppressesOutput() bool {
	return c.SuppressOutput
}

func (c PullCommand) GetWorkingDir() string {
	return "."
}
