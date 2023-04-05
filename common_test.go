package utilities

import (
	//"github.com/HimanshuQA87/gin/utilities"
	. "github.com/onsi/ginkgo/v2"
	//. "github.com/onsi/gomega"
)

var _ = Describe("Test one", func() {
	headers := CommonHeaders()
	FireGetRequest(headers, "https://simple-grocery-store-api.glitch.me/status")
})
