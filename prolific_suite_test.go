package main_test

import (
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

var _ = BeforeSuite(func() {
	var err error
	prolific, err = gexec.Build("github.com/onsi/prolific")
	Ω(err).ShouldNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
