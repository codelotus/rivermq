// +build integration

package route_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/codelotus/rivermq/model"
	. "github.com/codelotus/rivermq/route"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	err := CreateRiverMQDB()
	Expect(err).NotTo(HaveOccurred())
})

var _ = Describe("RoutesAndHandlers", func() {

	var validSub string
	var responseSub string

	BeforeEach(func() {
		validSub = `{"type":"msgType","callbackUrl":"http://localhost/endpoint"}`
		responseSub = `{"timestamp":"","type":"msgType","callbackUrl":"http://localhost/endpoint"}`
	})

	AfterEach(func() {
		_, err := QueryDB("DROP MEASUREMENT \"Subscription\"")
		if err != nil {
			Fail(err.Error())
		}
	})

	Describe("RoutingSubscriptions", func() {
		Context("Subscriptions Endpoints", func() {

			It("should route to CreateSubscriptionHandler", func() {
				buf := bytes.NewBufferString(validSub)
				req, _ := http.NewRequest("POST", "/subscriptions", buf)
				res := httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusCreated))
				var sub *Subscription
				json.NewDecoder(res.Body).Decode(&sub)
				Expect(sub.ID).ToNot(BeNil())
				Expect(sub.Timestamp).ToNot(BeNil())
			})

			It("should route to FindAllSubscriptionsHandler", func() {
				buf := bytes.NewBufferString(validSub)
				req, _ := http.NewRequest("POST", "/subscriptions", buf)
				res := httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusCreated))
				var sub *Subscription
				json.NewDecoder(res.Body).Decode(&sub)
				Expect(sub.ID).ToNot(BeNil())

				req, _ = http.NewRequest("GET", "/subscriptions", nil)
				res = httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				var subs *[]Subscription
				json.NewDecoder(res.Body).Decode(&subs)
				Expect(len(*subs)).To(Equal(1))
			})

			It("should route to FindAllSubscriptionsHandler and find by Subscription Type", func() {
				buf := bytes.NewBufferString(validSub)
				req, _ := http.NewRequest("POST", "/subscriptions", buf)
				res := httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusCreated))

				buf = bytes.NewBufferString(`{"type":"thing","callbackUrl":"http://localhost/endpoint"}`)
				req, _ = http.NewRequest("POST", "/subscriptions", buf)
				res = httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusCreated))
				var sub *Subscription
				json.NewDecoder(res.Body).Decode(&sub)
				Expect(sub.ID).ToNot(BeNil())

				req, _ = http.NewRequest("GET", "/subscriptions?type=thing", nil)
				res = httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				var subs *[]Subscription
				json.NewDecoder(res.Body).Decode(&subs)
				Expect(len(*subs)).To(Equal(1))
			})

			It("should route to FindSubscriptionByIDHandler", func() {
				buf := bytes.NewBufferString(validSub)
				req, _ := http.NewRequest("POST", "/subscriptions", buf)
				res := httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusCreated))
				var sub *Subscription
				json.NewDecoder(res.Body).Decode(&sub)
				Expect(sub.ID).ToNot(BeNil())

				url := "/subscriptions/" + sub.ID.String()
				req, _ = http.NewRequest("GET", url, nil)
				res = httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))
				var resultSub *Subscription
				json.NewDecoder(res.Body).Decode(&resultSub)
				Expect(resultSub.ID).To(BeEquivalentTo(sub.ID))
				Expect(resultSub.Timestamp).ToNot(BeNil())
				Expect(resultSub.Type).To(BeEquivalentTo(sub.Type))
				Expect(resultSub.CallbackURL).To(BeEquivalentTo(sub.CallbackURL))
			})

			It("should route to DeleteSubscriptionByIDHandler", func() {
				buf := bytes.NewBufferString(validSub)
				req, _ := http.NewRequest("POST", "/subscriptions", buf)
				res := httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusCreated))
				var sub *Subscription
				json.NewDecoder(res.Body).Decode(&sub)
				Expect(sub.ID).ToNot(BeNil())

				req, _ = http.NewRequest("DELETE", "/subscriptions/"+sub.ID.String(), nil)
				res = httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))

				req, _ = http.NewRequest("GET", "/subscriptions/"+sub.ID.String(), nil)
				res = httptest.NewRecorder()
				NewRiverMQRouter().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusNotFound))
			})

		})
	})
})
