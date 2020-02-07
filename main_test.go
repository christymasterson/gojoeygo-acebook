package main

import (
  "io/ioutil"
  "net/http"
  "net/http/httptest"
  "strings"
  "testing"
)

func TestShowHomePageUnauthenticated(t *testing.T) {

	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: actual %v expected %v", status, http.StatusOK)
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Home Page</title>") < 0 {
			t.Fail()
	}
}
