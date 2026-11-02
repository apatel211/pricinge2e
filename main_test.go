package main

import (
	"fmt"
	"os"
	"testing"
)

var TargetBaseURL string

func TestMain(m *testing.M) {
	fmt.Println("=== [START] Initializing Pricing E2E Test Suite ===")

	TargetBaseURL = os.Getenv("API_BASE_URL")
	if TargetBaseURL == "" {
		TargetBaseURL = "https://staging-api.pricing.local"
	}

	fmt.Printf("Configured Target Environment: %s\n", TargetBaseURL)

	exitCode := m.Run()

	fmt.Println("=== [END] Tearing Down Pricing E2E Test Suite ===")
	os.Exit(exitCode)
}