package model

// RideRequest represents the incoming request payload for pricing calculations
type RideRequest struct {
	ID        string  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	BaseFare  float64 `json:"base_fare"`
}

// PricingResponse represents the final calculated fare returned by the API
type PricingResponse struct {
	ID         string  `json:"id"`
	Multiplier float64 `json:"multiplier"`
	FinalFare  float64 `json:"final_fare"`
	IsSurge    bool    `json:"is_surge"`
}