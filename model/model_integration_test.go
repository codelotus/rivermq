// +build integration

package model_test

import (
	"math/rand"
	"time"

	. "github.com/codelotus/rivermq/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = BeforeSuite(func() {
	err := CreateRiverMQDB()
	Expect(err).NotTo(HaveOccurred())
})

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

	AfterEach(func() {
		_, err := QueryDB("DROP MEASUREMENT \"Subscription\"")
		if err != nil {
			Fail(err.Error())
		}
	})

	Describe("Saving a Subscription", func() {
		Context("Some Context", func() {
			It("should save a valid Subscription", func() {
				sub := Subscription{
					Type:        randStr(10),
					CallbackURL: "http://" + randStr(8) + "/endpoint",
				}
				err := SaveSubscription(sub)
				Expect(err).To(BeNil())
			})
			Measure("it should save subscriptions efficiently", func(b Benchmarker) {
				runtime := b.Time("runtime", func() {
					sub := Subscription{
						Type:        randStr(10),
						CallbackURL: "http://" + randStr(8) + "/endpoint",
					}
					err := SaveSubscription(sub)
					Expect(err).To(BeNil())
				})
				Expect(runtime.Seconds()).To(BeNumerically("<", 1.2), "SaveSubscription() shouldn't take too long.")
			}, 10)
		})
	})

	Describe("Reading Subscriptions", func() {
		Context("Some Context", func() {
			It("should find five subscriptions when five are saved", func() {
				for x := 0; x < 5; x++ {
					sub := Subscription{
						Type:        randStr(10),
						CallbackURL: "http://" + randStr(8) + "/endpoint",
					}
					err := SaveSubscription(sub)
					Expect(err).To(BeNil())
				}

				subs, err := GetAllSubscriptions()
				Expect(err).To(BeNil())
				Expect(len(subs)).To(BeEquivalentTo(5))
				/*
					fmt.Printf("len(res):\t%v\n", len(res))
					fmt.Printf("len(res[0].Series):\t%v\n", len(res[0].Series))
					fmt.Printf("res:\t%v\n", res)
					Expect(len(res[0].Series[0].Values)).To(BeEquivalentTo(5))

					fmt.Printf("Res Name:\t%v\n", res[0].Series[0].Name)
					fmt.Printf("Res Tags:\t%v\n", res[0].Series[0].Tags)
					fmt.Printf("Res Columns:\t%v\n", res[0].Series[0].Columns)
					fmt.Printf("Res Values[0]:\t%v\n", res[0].Series[0].Values[0])
					fmt.Printf("Res Value[0][0]:\t%v\n", res[0].Series[0].Values[0][0])
					fmt.Printf("len(Values):\t%v\n", len(res[0].Series[0].Values))
				*/
			})
		})
	})
})

// http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
var src = rand.NewSource(time.Now().UnixNano())

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func randStr(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
