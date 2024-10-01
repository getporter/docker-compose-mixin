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
	Subcommand       *Command
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
	// Final Command: docker-compose ARGS FLAGS [CMD CMD_ARGS CMD_FLAGS CMD_SUFFIX_ARGS]
	// We need to return: CMD CMD_ARGS CMD_FLAGS CMD_SUFFIX_ARGS
	if c.Subcommand == nil {
		return nil
	}

	cmd := (*c.Subcommand).GetCommand()
	cmdArgs := (*c.Subcommand).GetArguments()
	cmdFlags := (*c.Subcommand).GetFlags()
	cmdSuffixArgs := (*c.Subcommand).GetSuffixArguments()

	suffixArgs := make([]string, 0, 1+len(cmdArgs)+(len(cmdFlags)*2)+len(cmdSuffixArgs))
	suffixArgs = append(suffixArgs, cmd)
	suffixArgs = append(suffixArgs, cmdArgs...)
	suffixArgs = append(suffixArgs, cmdFlags.ToSlice(builder.DefaultFlagDashes)...)
	suffixArgs = append(suffixArgs, cmdSuffixArgs...)

	return suffixArgs
}

func (c ComposeCommand) GetOutputs() []builder.Output {
	// Go doesn't have generics, nothing to see here...
	outputs := make([]builder.Output, len(c.Outputs))
	for i := range c.Outputs {
		outputs[i] = c.Outputs[i]
	}
	if c.Subcommand != nil {
		outputs = append(outputs, (*c.Subcommand).GetOutputs()...)
	}
	return outputs
}

func (c ComposeCommand) SuppressesOutput() bool {
	if c.Subcommand != nil && (*c.Subcommand).SuppressesOutput() {
		return true
	} else {
		return c.SuppressOutput
	}
}

func (c ComposeCommand) GetWorkingDir() string {
	return c.WorkingDirectory
}
