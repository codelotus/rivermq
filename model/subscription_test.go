package model

import (
	"encoding/json"

	"github.com/pborman/uuid"

	client "github.com/influxdata/influxdb/client/v2"
	"github.com/influxdata/influxdb/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subscription", func() {

	var (
		validSub   Subscription
		invalidSub Subscription
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
	})

	Describe("Validating a Subscription", func() {
		Context("With valid values", func() {
			It("should validate", func() {
				Expect(ValidateSubscription(validSub)).To(BeTrue())
			})
		})
		Context("With invalidValues", func() {
			It("should not validate", func() {
				res, err := ValidateSubscription(invalidSub)
				Expect(res).ToNot(BeTrue())
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Converting a Influx Client Result to a Subscription", func() {
		Context("Some context", func() {
			It("should successfully convert", func() {
				mockResult := createMockInfluxResult()
				res, err := convertResultToSubscriptionSlice(mockResult)
				Expect(err).To(BeNil())
				Expect(res).NotTo(BeNil())
			})
		})
		Measure("it should convert influxdb results efficiently", func(b Benchmarker) {
			runtime := b.Time("runtime", func() {
				mockResult := createMockInfluxResult()
				res, err := convertResultToSubscriptionSlice(mockResult)
				Expect(err).To(BeNil())
				Expect(res).NotTo(BeNil())
			})
			Expect(runtime.Seconds()).To(BeNumerically("<", 0.1), "ConvertResultToSubscriptionSlice() is to slow")
		}, 1000)
	})
})

func createMockInfluxResult() (res []client.Result) {
	var values [][]interface{}
	id := uuid.NewUUID().String()
	values = append(values, []interface{}{id, json.Number(123123123123123), "messageType", "http://localhost/endpoint"})

	seriesSlice := []models.Row{}
	seriesSlice = append(seriesSlice, models.Row{
		Name:    "Subscription",
		Tags:    make(map[string]string, 0),
		Columns: []string{"ID", "time", "Type", "CallbackURL"},
		Values:  values,
	})

	result := []client.Result{}
	result = append(result, client.Result{
		Series: seriesSlice,
	})

	return result
}
