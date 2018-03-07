package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// method, route, request_body
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	// Recorder acts as target for http request (aka sudo browser)
	recorder := httptest.NewRecorder()

	// handler is the function in main.go we want to test
	hf := http.HandlerFunc(handler)

	// serve the HTTP request to records (thus executes handler)
	hf.ServeHTTP(recorder, req)

	// Check that we get the correct status code
	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// check the response body is as expected
	expected := "Hello World!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
