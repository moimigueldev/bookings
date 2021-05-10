package main

import (
	"net/http"
	"os"
	"testing"
)

// Before you start running the test do something inside this function, then run the test, then exit

func testMain(m *testing.M) {

	os.Exit(m.Run())
}

type myHandler struct {
}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
