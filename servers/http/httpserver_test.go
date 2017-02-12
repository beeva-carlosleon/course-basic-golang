package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func ResponseCreator() *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", "http://example.com/hello", nil)
	w := httptest.NewRecorder()
	HelloServer(w, req)
	fmt.Printf("%d - %s", w.Code, w.Body.String())
	return w
}

func testSimpleChecker(result, expected interface{}, t *testing.T) {
	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}

func TestHandler(t *testing.T) {
	resp := ResponseCreator()
	testSimpleChecker(resp.Code, 200, t)
	testSimpleChecker(resp.Body.String(), responseBody, t)
}
func TestServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HelloServer))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	testSimpleChecker(string(greeting), responseBody, t)
}

func TestLaunchServer(t *testing.T) {
	port := 12345
	go launchServer(&port)
	time.Sleep(time.Second)
	resp, err := http.Get(fmt.Sprintf("http://localhost:%d/hello", 12345))
	if err != nil {
		t.Errorf("UnExpected error %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	testSimpleChecker(string(body), responseBody, t)
}
