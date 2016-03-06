package model_test

import (
	"github.com/codelotus/rivermq/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

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
