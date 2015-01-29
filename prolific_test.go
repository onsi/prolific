package main_test

import (
	"bytes"
	"encoding/csv"
	"os"
	"os/exec"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Prolific", func() {
	var session *gexec.Session
	var err error

	AfterEach(func() {
		os.Remove("stories.prolific")
	})

	Describe("prolific template", func() {
		BeforeEach(func() {
			cmd := exec.Command(prolific, "template")
			session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should generate a template file", func() {
			_, err := os.Stat("stories.prolific")
			Ω(err).ShouldNot(HaveOccurred())
		})
	})

	Describe("generating csv output", func() {
		BeforeEach(func() {
			cmd := exec.Command(prolific, "template")
			session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))

			cmd = exec.Command(prolific, "Onsi Fakhouri", "stories.prolific")
			session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
			Ω(err).ShouldNot(HaveOccurred())
			Eventually(session).Should(gexec.Exit(0))
		})

		It("should convert the passed-in prolific file, honoring the passed-in author", func() {
			reader := csv.NewReader(bytes.NewReader(session.Out.Contents()))
			records, err := reader.ReadAll()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(records[0]).Should(Equal([]string{"Requested By", "Title", "Description", "Labels"}))
			Ω(records).Should(HaveLen(4))
			Ω(records[1][0]).Should(Equal("Onsi Fakhouri"))
			Ω(records[1][1]).Should(Equal("As a user I can toast a bagel"))
			Ω(records[1][2]).Should(Equal("When I insert a bagel into toaster and press the on button, I should get a toasted bagel"))
			Ω(records[1][3]).Should(Equal("mvp,toasting"))
			Ω(records[3][3]).Should(Equal("mvp,clean-up"))
		})
	})
})
