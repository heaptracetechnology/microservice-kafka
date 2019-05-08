package kafka

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Kafka", func() {


	var sub Subscribe

	sub.Topic ="hello"

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