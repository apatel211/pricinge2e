package main

import (
	"pricinge2e/model"
	"testing"
)

func TestFareComparisons(t *testing.T) {
	testCases := []struct {
		name          string
		inputRequest  model.RideRequest
		multiplier    float64
		expectedTotal float64
	}{
		{
			name:         "Standard Off-peak Calculation",
			inputRequest: model.RideRequest{BaseFare: 10.00},
			multiplier:   1.0,
			expectedTotal: 10.00,
		},
		{
			name:         "Peak Weekend Surge Multiplier",
			inputRequest: model.RideRequest{BaseFare: 20.00},
			multiplier:   2.0,
			expectedTotal: 40.00,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculatedTotal := tc.inputRequest.BaseFare * tc.multiplier
			if calculatedTotal != tc.expectedTotal {
				t.Errorf("%s failed -> Expected $%.2f, but computed $%.2f", tc.name, tc.expectedTotal, calculatedTotal)
			}
		})
	}
}