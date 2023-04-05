package utilities

import (
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/gomega"
)

func FireGetRequest(headers map[string]string, endpoint string) {

	client := resty.New().SetHeaders(headers)
	resp, _ := client.R().Get(endpoint)
	Expect(resp.StatusCode()).To(Equal(200))

}

func CommonHeaders() map[string]string {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	return headers
}
