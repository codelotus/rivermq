package model_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	model "github.com/codelotus/rivermq/model"
	"testing"
)

func TestModel(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Model Suite")
}

var _ = BeforeSuite(func() {
	err := model.CreateRiverMQDB()
	Expect(err).NotTo(HaveOccurred())
})
