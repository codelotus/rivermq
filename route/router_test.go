package route_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "github.com/codelotus/rivermq/route"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RoutesAndHandlers", func() {

	var validSub string
	var responseSub string

	BeforeEach(func() {
		validSub = `{"type":"msgType","callbackUrl":"http://localhost/endpoint"}
`
		responseSub = `{"timestamp":"","type":"msgType","callbackUrl":"http://localhost/endpoint"}
`
	})

	Describe("RoutingSubscriptions", func() {
		Context("Subscriptions Endpoints", func() {

			It("should route to CreateSubscriptionHandler", func() {
				buf := bytes.NewBufferString(validSub)
				req, _ := http.NewRequest("POST", "/subscriptions", buf)
				res := httptest.NewRecorder()
				NewRouter().ServeHTTP(res, req)
				Expect(res.Body.String()).To(Equal(responseSub),
					"response[%s] does not match %s", res.Body.String(), validSub)
			})

			It("should route to GetSubscriptionByIDHandler", func() {
				req, _ := http.NewRequest("GET", "/subscriptions/12", nil)
				res := httptest.NewRecorder()
				NewRouter().ServeHTTP(res, req)
				Expect(res.Body.String()).To(Equal("GetSubscription: 12"))
			})

			It("should route to GetAllSubscriptionsHandler", func() {
				req, _ := http.NewRequest("GET", "/subscriptions", nil)
				res := httptest.NewRecorder()
				NewRouter().ServeHTTP(res, req)
				Expect(res.Body.String()).To(Equal("GetAllSubscriptions"))
			})

			It("should route to DeleteSubscriptionByIDHandler", func() {
				req, _ := http.NewRequest("DELETE", "/subscriptions/99", nil)
				res := httptest.NewRecorder()
				NewRouter().ServeHTTP(res, req)
				Expect(res.Body.String()).To(Equal("DeleteSubscriptionByID: 99"))
			})

		})
	})
})
