package config_test

import (
	"io/ioutil"
	"os"

	"github.com/ankeesler/flexec/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var storeDir string

	BeforeEach(func() {
		var err error
		storeDir, err = ioutil.TempDir("", "flexec_config_test")
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(os.RemoveAll(storeDir)).To(Succeed())
	})

	It("writes and reads a config", func() {
		cWrite := &config.Config{
			ParamDefaults: map[string]string{
				"tuna":   "fish",
				"marlin": "foo",
			},
		}
		Expect(config.Write(storeDir, cWrite)).To(Succeed())

		cRead, err := config.Read(storeDir)
		Expect(err).NotTo(HaveOccurred())
		Expect(cRead).To(Equal(cWrite))
	})
})
