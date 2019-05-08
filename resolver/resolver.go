package resolver

import (
	"github.com/concourse/concourse/atc"
)

//go:generate counterfeiter . Resolver
//go:generate counterfeiter . Handler

type Resolver interface {
	ResolveInput(name string) (string, error)
	ResolveOutput(name string) (string, error)
	ResolveParam(paramKey string) (string, error)
}

type Handler interface {
	OnInput(name, path string)
	OnOutput(name, path string)
	OnParam(paramKey, paramValue string)
}

type Runner struct {
	handler   Handler
	resolvers []Resolver
}

func NewRunner(handler Handler, resolvers ...Resolver) *Runner {
	return &Runner{
		handler:   handler,
		resolvers: resolvers,
	}
}

func (r *Runner) Run(task *atc.TaskConfig) error {
	return nil
}
