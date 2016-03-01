package model_test

import (
	. "github.com/codelotus/rivermq/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Model", func() {

	var (
		validSub Subscription
	)

	BeforeEach(func() {
		validSub = Subscription{
			Type:        "subscriptionType",
			CallbackURL: "http://localhost/abc",
		}
	})

	Describe("Saving a Subscription", func() {
		Context("Some Context", func() {
			It("should save a valid Subscription", func() {
				err := SaveSubscription(validSub)
				Expect(err).To(BeNil())
			})
			Measure("it should save subscriptions efficiently", func(b Benchmarker) {
				runtime := b.Time("runtime", func() {
					err := SaveSubscription(validSub)
					Expect(err).To(BeNil())
				})
				Expect(runtime.Seconds()).To(BeNumerically("<", 0.1), "SaveSubscription() shouldn't take too long.")
			}, 10)
		})
	})
})
