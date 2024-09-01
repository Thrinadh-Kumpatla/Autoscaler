package main

import (
	"log"

	"github.com/Thrinadh-Kumpatla/auto-scaler/internal/config"
	"github.com/Thrinadh-Kumpatla/auto-scaler/internal/scaler"
)

func main() {
	cfg := config.Parse()
	log.Printf("Starting auto-scaler for application at %s", cfg.BaseURL)
	
	s := scaler.New(cfg)
	s.Run()
}