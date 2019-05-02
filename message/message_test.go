package message

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cron event subscribe", func() {
	var argumentData Subscribe
	argumentData.IsTesting = true
	var data Data
	data.Interval = 2
	data.InitialDelay = 1
	argumentData.Data = data
	argumentData.Endpoint = "http://webhook.site/bfd1aea6-0562-4087-90a3-68efab7d0302"
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(argumentData)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(TriggerCron)
	handler.ServeHTTP(recorder, request)

	Describe("Send Message", func() {
		Context("SendMessage", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
