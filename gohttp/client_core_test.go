package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {
	// Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.builder.headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Errorf("We excpect 3 headers")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type")
	}
	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid content type received")
	}
}

func TestGetRequestBody(t *testing.T) {
	// Initialization
	client := httpClient{}

	t.Run("NoBodyNilRespone", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)
		if err != nil {
			t.Error("no error expected when passing nil body")
		}

		if body != nil {
			t.Error("no error expected when passing nil body")
		}
	})

	t.Run("BodyWithJSON", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})

	t.Run("BodyWithXML", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/xml", requestBody)
		if err != nil {
			t.Error("no error expected when marshaling slice to xml")
		}

		if string(body) != `<string>one</string><string>two</string>` {
			t.Error("failed to get expected xml format")
		}
	})

	t.Run("BodyWithJSONAsDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/text", requestBody)
		if err != nil {
			t.Error("no error expected when marshaling slice as json")
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}
	})
}
