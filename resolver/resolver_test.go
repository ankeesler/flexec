package resolver_test

import (
	"github.com/ankeesler/flexec/resolver"
	"github.com/ankeesler/flexec/resolver/resolverfakes"
	"github.com/concourse/concourse/atc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resolver", func() {
	var (
		handler   *resolverfakes.FakeHandler
		resolverA *resolverfakes.FakeResolver
		resolverB *resolverfakes.FakeResolver
		runner    *resolver.Runner
	)

	BeforeEach(func() {
		handler = &resolverfakes.FakeHandler{}
		resolverA = &resolverfakes.FakeResolver{}
		resolverB = &resolverfakes.FakeResolver{}
		runner = resolver.NewRunner(handler, resolverA, resolverB)
	})

	It("calls the first resolver for each input, output, param", func() {
		resolverA.ResolveInputReturnsOnCall(0, "some-resolved-input", nil)
		resolverA.ResolveInputReturnsOnCall(1, "some-other-resolved-input", nil)

		resolverA.ResolveOutputReturnsOnCall(0, "some-resolved-output", nil)
		resolverA.ResolveOutputReturnsOnCall(1, "some-other-resolved-output", nil)

		resolverA.ResolveParamReturnsOnCall(0, "fish", nil)
		resolverA.ResolveParamReturnsOnCall(1, "marlin", nil)

		task := atc.TaskConfig{
			Inputs: []atc.TaskInputConfig{
				atc.TaskInputConfig{
					Name: "some-input",
				},
				atc.TaskInputConfig{
					Name: "some-other-input",
				},
			},
			Outputs: []atc.TaskOutputConfig{
				atc.TaskOutputConfig{
					Name: "some-output",
				},
				atc.TaskOutputConfig{
					Name: "some-other-output",
				},
			},
			Params: map[string]string{
				"tuna":   "fish",
				"marlin": "bass",
			},
		}
		Expect(runner.Run(&task)).To(Succeed())

		Expect(resolverA.ResolveInputCallCount()).To(Equal(2))
		Expect(resolverA.ResolveInputArgsForCall(0)).To(Equal("some-input"))
		Expect(resolverA.ResolveInputArgsForCall(1)).To(Equal("some-other-input"))
		Expect(resolverA.ResolveOutputCallCount()).To(Equal(2))
		Expect(resolverA.ResolveOutputArgsForCall(0)).To(Equal("some-output"))
		Expect(resolverA.ResolveOutputArgsForCall(1)).To(Equal("some-other-output"))
		Expect(resolverA.ResolveParamCallCount()).To(Equal(2))
		Expect(resolverA.ResolveParamArgsForCall(0)).To(Equal("tuna"))
		Expect(resolverA.ResolveParamArgsForCall(1)).To(Equal("marlin"))

		Expect(resolverB.ResolveInputCallCount()).To(Equal(0))
		Expect(resolverB.ResolveOutputCallCount()).To(Equal(0))
		Expect(resolverB.ResolveParamCallCount()).To(Equal(0))

		Expect(handler.OnInputCallCount()).To(Equal(1))
		name, path := handler.OnInputArgsForCall(0)
		Expect(name).To(Equal("some-input"))
		Expect(path).To(Equal("some-resolved-input"))
	})
})
