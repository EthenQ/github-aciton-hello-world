package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func testRequest(t *testing.T, ts *httptest.Server, method, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}

func TestHelloHandler(t *testing.T) {

	r := chi.NewRouter()
	r.Method("GET", "/", Handler(helloHandler))
	ts := httptest.NewServer(r)
	defer ts.Close()
	if _, body := testRequest(t, ts, "GET", "/", nil); body != "hello world" {
		t.Fatalf(body)
	}
}
