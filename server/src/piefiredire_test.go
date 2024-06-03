package piefiredire

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetBeefSummary(t *testing.T) {
	mockResponse := "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone."

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, mockResponse)
	}))
	defer ts.Close()

	originalHTTPGet := httpGet
	defer func() { httpGet = originalHTTPGet }()

	httpGet = func(url string) (*http.Response, error) {
		return ts.Client().Get(ts.URL)
	}

	result := GetBeefSummary()

	expected := map[string]int32{
		"t-bone":   4,
		"fatback":  1,
		"pastrami": 1,
		"pork":     1,
		"meatloaf": 1,
		"jowl":     1,
		"enim":     1,
		"bresaola": 1,
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("unexpected result: got %v want %v", result, expected)
	}
}
