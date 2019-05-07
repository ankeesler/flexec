package task_test

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/concourse/concourse/atc"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	yaml "gopkg.in/yaml.v1"
)

var _ = Describe("Task", func() {
	var taskConfig *os.File

	BeforeEach(func() {
		var err error
		taskConfig, err = ioutil.TempFile("", "flexec_task_test")
		Expect(err).To(Succeed())

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
		taskData, err := yaml.Marshal(&task)
		Expect(err).To(Succeed())
		fmt.Fprintln(taskConfig, string(taskData))
	})

	AfterEach(func() {
		Expect(taskConfig.Close()).To(Succeed())
		Expect(os.RemoveAll(taskConfig.Name())).To(Succeed())
	})

	It("reads a concourse task.yml file", func() {
	})
})
