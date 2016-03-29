package model

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Message", func() {

	var (
		validMsg   Message
		invalidMsg Message
	)

	BeforeEach(func() {
		validMsg = Message{
			Type: "messageType",
			Body: []byte("{}"),
		}
		invalidMsg = Message{
			Body: []byte("{}"),
		}
	})

	Describe("Validating a Message", func() {
		Context("with valid values", func() {
			It("should validate", func() {
				Expect(ValidateMessage(validMsg)).To(BeTrue())
			})
		})
		Context("with invalid values", func() {
			It("should not validate", func() {
				res, err := ValidateMessage(invalidMsg)
				Expect(res).To(BeFalse())
				Expect(err).ToNot(BeNil())
			})
		})
	})

})
