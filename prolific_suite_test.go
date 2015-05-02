package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"

	"testing"
)

func TestProlific(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prolific Suite")
}

var prolific string
var workingDir string

var _ = BeforeSuite(func() {
	var err error
	prolific, err = gexec.Build("github.com/onsi/prolific")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = BeforeEach(func() {
	var err error
	workingDir, err = ioutil.TempDir("", "prolific")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterEach(func() {
	Ω(os.RemoveAll(workingDir)).Should(Succeed())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

func Command(path string, args ...string) *exec.Cmd {
	cmd := exec.Command(path, args...)
	cmd.Dir = workingDir
	return cmd
}
