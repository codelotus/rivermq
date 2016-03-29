// +build integration

package model_test

import (
	. "github.com/codelotus/rivermq/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {

	var (
		validMessage Message
	)

	AfterEach(func() {
		_, err := QueryDB("DROP MEASUREMENT \"Message\"")
		if err != nil {
			Fail(err.Error())
		}
	})

	Describe("Message", func() {
		Context("Saving", func() {
			It("should save a valid Message", func() {
				validMessage.Type = RandomString(10)
				validMessage.Body = []byte("{ \"prop:\", \"key\"}")
				_, err := SaveMessage(validMessage)
				Expect(err).To(BeNil())
			})
			Measure("it should save messages efficiently", func(b Benchmarker) {
				runtime := b.Time("runtime", func() {
					validMessage.Type = RandomString(10)
					validMessage.Body = []byte("{ \"prop:\", \"key\"}")
					_, err := SaveMessage(validMessage)
					Expect(err).To(BeNil())
				})
				Expect(runtime.Seconds()).To(BeNumerically("<", 1.2), "SaveMessage should be efficient")
			}, 10)
		})
	})
})
