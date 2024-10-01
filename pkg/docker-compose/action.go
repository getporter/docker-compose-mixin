package dockercompose

import (
	"get.porter.sh/mixin/docker-compose/pkg/docker-compose/commands"
	"get.porter.sh/porter/pkg/exec/builder"
	"gopkg.in/yaml.v3"
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
	commands.ComposeCommand `yaml:"docker-compose"`
}

// UnmarshalYAML takes any yaml in this form
//
//	docker-compose:
//	  description: something
//	  COMMAND: # e.g. pull/up/down -> make the PullCommand/UpCommand/DownCommand for us
func (s *Step) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Turn the yaml into a raw map so we can iterate over the values and
	// look for which command was used
	stepMap := map[string]map[string]interface{}{}
	err := unmarshal(&stepMap)
	if err != nil {
		return err
	}

	// Get at the values defined under "docker-compose"
	composeStep := stepMap["docker-compose"]
	b, err := yaml.Marshal(composeStep)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(b, &s.ComposeCommand)
	if err != nil {
		return err
	}

	// Turn the command into its typed data structure
	for key, value := range composeStep {
		var cmd commands.Command

		switch key {
		case "down":
			cmd = &commands.DownCommand{}
		case "pull":
			cmd = &commands.PullCommand{}
		case "up":
			cmd = &commands.UpCommand{}
		default:
			continue
		}

		b, err = yaml.Marshal(value)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(b, cmd)
		if err != nil {
			return err
		}

		s.Subcommand = &cmd
		break // There is only 1 command
	}

	return nil
}
