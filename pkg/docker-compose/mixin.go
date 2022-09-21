package dockercompose

import "get.porter.sh/porter/pkg/runtime"

type Mixin struct {
	runtime.RuntimeConfig
	// add whatever other context/state is needed here
}

// New docker-compose mixin client, initialized with useful defaults.
func New() *Mixin {
	return &Mixin{
		RuntimeConfig: runtime.NewConfig(),
	}
}
