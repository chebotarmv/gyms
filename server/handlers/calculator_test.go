package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	req, err := http.NewRequest("GET", "/calculate?item=10", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CalculatePacks)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %v but got %v", http.StatusOK, status)
	}

	var respMap map[int]int
	err = json.Unmarshal(rr.Body.Bytes(), &respMap)
	if err != nil {
		t.Errorf("Error decoding response body: %v", err)
	}

	expectedResp := map[int]int{250: 1}

	if !mapsEqual(respMap, expectedResp) {
		t.Errorf("Response map does not match expected values. Got: %v, Expected: %v", respMap, expectedResp)
	}
}

func mapsEqual(a, b map[int]int) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok || valueA != valueB {
			return false
		}
	}

	return true
}
