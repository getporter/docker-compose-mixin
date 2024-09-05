package commands

import "get.porter.sh/porter/pkg/exec/builder"

var _ Command = ComposeCommand{}

type ComposeCommand struct {
	Description      string        `yaml:"description"`
	WorkingDirectory string        `yaml:"dir,omitempty"`
	Arguments        []string      `yaml:"arguments,omitempty"`
	Flags            builder.Flags `yaml:"flags,omitempty"`
	Outputs          []Output      `yaml:"outputs,omitempty"`
	SuppressOutput   bool          `yaml:"suppress-output,omitempty"`
}

func (c ComposeCommand) GetCommand() string {
	return "docker-compose"
}

func (c ComposeCommand) GetArguments() []string {
	return c.Arguments
}

func (c ComposeCommand) GetFlags() builder.Flags {
	return c.Flags
}

func (c ComposeCommand) GetSuffixArguments() []string {
	return nil
}

func (c ComposeCommand) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(c.Outputs))
	for i := range c.Outputs {
		outputs[i] = c.Outputs[i]
	}
	return outputs
}

func (c ComposeCommand) SuppressesOutput() bool {
	return c.SuppressOutput
}

func (c ComposeCommand) GetWorkingDir() string {
	return c.WorkingDirectory
}
