package main

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAliveHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/status/alive", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AliveHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got '%v' want '%v'",
			status, http.StatusOK)
	}

	want := statusResponse{
		Status: "Namer service is alive",
	}

	got := statusResponse{}
	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	// if expected.Status != response.Status {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("healthyHandler mismatch (-want +got):\n%s", diff)
	}
}

func TestHealthyHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/status/heathy", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthyHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got '%v' want '%v'",
			status, http.StatusOK)
	}

	want := statusResponse{
		Status: "Namer service is healthy",
	}

	got := statusResponse{}
	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatal(err)
	}

	// if expected.Status != response.Status {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("healthyHandler mismatch (-want +got):\n%s", diff)
	}
}
