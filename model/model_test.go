package model_test

import (
	. "github.com/codelotus/rivermq/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Model", func() {

	var (
		validSub   Subscription
		invalidSub Subscription
		message    Message
	)

	BeforeEach(func() {
		validSub = Subscription{
			Type:        "messageType",
			CallbackURL: "http://localhost:1234/msg",
		}

		invalidSub = Subscription{
			Type:        "messageType",
			CallbackURL: "http//localhost:1234/msg",
		}

		message = Message{
			Type: "messageType",
			Body: "message body",
		}
	})

	Describe("Validating a Subscription", func() {
		Context("With valid values", func() {
			It("should validate", func() {
				Expect(validSub.Validate()).To(BeTrue())
			})
		})
		Context("With invalidValues", func() {
			It("should not validate", func() {
				result, error := invalidSub.Validate()
				Expect(result).ToNot(BeTrue())
				Expect(error).To(HaveOccurred())
			})
		})
	})

})
