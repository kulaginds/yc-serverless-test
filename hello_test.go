package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_helloHandler(t *testing.T) {
	webServer := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer webServer.Close()

	resp, err := http.Get(webServer.URL)
	if err != nil {
		t.Fatalf("error on http get: %v\n", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error on read response body: %v\n", err)
	}

	if bytes.Compare(respBody, helloMessage) != 0 {
		t.Fatalf("response body not equal hello message: %s\n", string(respBody))
	}
}
