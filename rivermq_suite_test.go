package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRivermq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rivermq Suite")
}
