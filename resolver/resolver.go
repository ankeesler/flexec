package resolver

type Resolver interface {
	ResolveInput(name string) (string, error)
	ResolveOutput(name string) (string, error)
	ResolveParam(paramKey string) (string, error)
}

type Runner struct {
	resolvers []Resolver
}

func NewRunner(resolvers ...Resolver) *Runner {
	return &Runner{
		resolvers: resolvers,
	}
}
