package kafka

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var host = os.Getenv("KAFKA_HOST")

var _ = Describe("Produce", func() {

	os.Setenv("HOST", host)

	produce := Produce{Topic: "hello", Message: "world"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(produce)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/produce", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ProduceStream)
	handler.ServeHTTP(recorder, request)

	Describe("Create topic if not exits and push message", func() {
		Context("Produce", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Kafka consume", func() {

	var sub Subscribe
	os.Setenv("HOST", host)
	sub.Data.Topic = "hello"
	sub.IsTesting = true
	sub.Endpoint = "http://webhook.site/bfd1aea6-0562-4087-90a3-68efab7d0302"
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(sub)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/consume", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Consume)
	handler.ServeHTTP(recorder, request)

	Describe("consume", func() {
		Context("consume", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})
