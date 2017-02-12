package main

import "testing"

func testSimpleChecker(result, expected interface{}, t *testing.T) {
	if result != expected {
		t.Errorf("Expected '%v', got '%v'", expected, result)
	}
}
func TestServer(t *testing.T) {
	server()
	testSimpleChecker(client(), nil, t)
}
