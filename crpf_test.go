package cprf_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/markelog/cprf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cprf", func() {
	var (
		from string
		to   string
	)

	BeforeEach(func() {
		from = "testdata/from"
		to = "testdata/to"
	})

	Describe("Throw when destination doesn't exist", func() {
		BeforeEach(func() {
			from += "/a/b/test"
			to += "/a/b/test"
		})

		It("should exist", func() {
			err := Copy(from, to)

			Ω(err).Should(HaveOccurred())
		})
	})

	Describe("Copy nested test file", func() {
		BeforeEach(func() {
			from += "/a/b/test"
			to += "/"
			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "test")
		})

		It("should exist", func() {
			data, _ := ioutil.ReadFile(to + "test")

			Ω(string(data)).To(Equal("1\n"))
		})
	})

	Describe("Override existing symlink in existing folder", func() {
		BeforeEach(func() {
			from += "/nested"
			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "/nested")
		})

		It("should exist", func() {
			err := Copy(from+"/symlink", to+"/nested")

			Ω(err).To(BeNil())
		})
	})

	Describe("Copy nested and override test file", func() {
		BeforeEach(func() {
			from += "/a/b/test"
			to += "/"

			hello := []byte("hello\ngo\n")
			ioutil.WriteFile(to+"test", hello, 0777)

			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "test")
		})

		It("should exist", func() {
			data, _ := ioutil.ReadFile(to + "test")

			Ω(string(data)).To(Equal("1\n"))
		})
	})

	Describe("Copy nested directory", func() {
		BeforeEach(func() {
			from += "/nested"

			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "/nested")
		})

		It("should copy symlink", func() {
			data, _ := os.Readlink(to + "/nested/symlink")

			Ω(data).To(Equal("test"))
		})

		It("should copy file and folder", func() {
			data, _ := ioutil.ReadFile(to + "/nested/a/b/test")

			Ω(string(data)).To(Equal("1\n"))
		})
	})

	Describe("Copy everything in the folder", func() {
		BeforeEach(func() {
			from += "/a/"
			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(filepath.Join(to + "/b"))
		})

		It("should exist", func() {
			data, _ := ioutil.ReadFile(to + "/b/test")

			Ω(string(data)).To(Equal("1\n"))
		})
	})

	Describe("Copy directory with hierarchy", func() {
		BeforeEach(func() {
			from += "/a"
			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(filepath.Join(to + "/a/b"))
		})

		It("should exist", func() {
			data, _ := ioutil.ReadFile(to + "/a/b/test")

			Ω(string(data)).To(Equal("1\n"))
		})
	})

	Describe("Override existing file", func() {
		BeforeEach(func() {
			from += "/test"

			hello := []byte("hello\ngo\n")
			ioutil.WriteFile(to+"/test", hello, 0777)

			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "/test")
		})

		It("should exist", func() {
			data, _ := ioutil.ReadFile(to + "/test")

			Ω(string(data)).To(Equal("1\n"))
		})
	})

	Describe("Copy symlink", func() {
		BeforeEach(func() {
			from += "/symlink"
			to += "/"

			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "symlink")
		})

		It("should exist", func() {
			fmt.Println(to + "symlink")
			data, _ := os.Readlink(to + "symlink")

			Ω(data).To(Equal("test"))
		})
	})

	Describe("Copy non-existent symlink", func() {
		BeforeEach(func() {
			from += "/non-existent-symlink"
			to += "/"

			Copy(from, to)
		})

		AfterEach(func() {
			os.RemoveAll(to + "non-existent-symlink")
		})

		It("should exist", func() {
			data, _ := os.Readlink(to + "non-existent-symlink")

			Ω(data).To(Equal("non-existent-symlink"))
		})
	})
})
