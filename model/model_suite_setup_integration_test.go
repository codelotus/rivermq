// +build integration

package model_test

import (
	"github.com/codelotus/rivermq/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	err := model.CreateRiverMQDB()
	Expect(err).NotTo(HaveOccurred())
})
