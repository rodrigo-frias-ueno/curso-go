package main

import (
	"log"

	"github.com/ueno-tecnologia-org/go-core/pkg/telemetry"
)

func buildTelemetryClient() telemetry.Client {
	telemetryConfig := telemetry.Config{
		ApplicationName: "categories-api",
		NewRelicLicense: "YOUR_NEW_RELIC_LICENSE_KEY",
	}

	telemetryClient, err := telemetry.NewClient(telemetryConfig)
	if err != nil {
		log.Fatal("Failed to create telemetry client:", err)
	}

	return telemetryClient
}
