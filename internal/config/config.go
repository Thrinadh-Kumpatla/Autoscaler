package config

import (
	"flag"
	"fmt"
	"time"
)

type Config struct {
	BaseURL       string
	TargetCPU     float64
	CheckInterval time.Duration
}

func Parse() *Config {
	port := flag.Int("port", 8123, "Port of the application")
	targetCPU := flag.Float64("target-cpu", 0.80, "Target CPU utilization (0.0-1.0)")
	checkInterval := flag.Duration("check-interval", 10*time.Second, "Interval between checks")
	flag.Parse()

	return &Config{
		BaseURL:       fmt.Sprintf("http://localhost:%d", *port),
		TargetCPU:     *targetCPU,
		CheckInterval: *checkInterval,
	}
}