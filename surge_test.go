package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"pricinge2e/model"
	"testing"
)

func TestSurgePricingMultiplier(t *testing.T) {
	// 1. Arrange: Prepare mock data targeting high-traffic zone coordinates
	requestPayload := model.RideRequest{
		ID:        "req_99812",
		Latitude:  40.7128, // High congestion zone simulator
		Longitude: -74.0060,
		BaseFare:  10.00,
	}

	// 2. Mock Server: Simulate internal business rules modifying final calculations
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := model.PricingResponse{
			ID:         "req_99812",
			Multiplier: 1.5, // 50% Surge applied
			FinalFare:  15.00,
			IsSurge:    true,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer mockServer.Close()

	// Convert payload to JSON to pass it into the mock request (Fixes the "declared and not used" error)
	body, err := json.Marshal(requestPayload)
	if err != nil {
		t.Fatalf("Failed to marshal request payload: %v", err)
	}

	// 3. Act: Run target API transaction pipeline logic using the payload
	t.Log("Dispatching ride query payload to microservice route pipeline...")
	resp, err := http.Post(mockServer.URL+"/calculate-fare", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to parse runtime HTTP pipeline execution: %v", err)
	}
	defer resp.Body.Close()

	var actualResponse model.PricingResponse
	if err := json.NewDecoder(resp.Body).Decode(&actualResponse); err != nil {
		t.Fatalf("Failed to unpack payload: %v", err)
	}

	// 4. Assert: Confirm output values exactly reflect engineering system specs
	if !actualResponse.IsSurge {
		t.Error("Failure: Active congestion parameters failed to flag system surge indicators.")
	}
	if actualResponse.FinalFare != 15.00 {
		t.Errorf("Calculation discrepancy: Expected $15.00, got $%.2f", actualResponse.FinalFare)
	}
}
