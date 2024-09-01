package scaler

import (
	"log"
	"time"

	"github.com/Thrinadh-Kumpatla/auto-scaler/internal/api"
	"github.com/Thrinadh-Kumpatla/auto-scaler/internal/config"
)

type Scaler struct {
	client        *api.Client
	targetCPU     float64
	checkInterval time.Duration
}

func New(cfg *config.Config) *Scaler {
	return &Scaler{
		client:        api.NewClient(cfg.BaseURL),
		targetCPU:     cfg.TargetCPU,
		checkInterval: cfg.CheckInterval,
	}
}

func (s *Scaler) Run() {
	for {
		status, err := s.client.GetStatus()
		if err != nil {
			log.Printf("Error getting status: %v", err)
			time.Sleep(s.checkInterval)
			continue
		}

		log.Printf("Current CPU: %.2f, Current Replicas: %d", status.CPU.HighPriority, status.Replicas)

		newReplicas := s.calculateNewReplicas(status.CPU.HighPriority, status.Replicas)

		if newReplicas != status.Replicas {
			log.Printf("Updating replicas to: %d", newReplicas)
			err := s.client.UpdateReplicas(newReplicas)
			if err != nil {
				log.Printf("Error updating replicas: %v", err)
			} else {
				log.Println("Successfully updated replicas")
			}
		} else {
			log.Println("No change in replicas needed")
		}

		time.Sleep(s.checkInterval)
	}
}

func (s *Scaler) calculateNewReplicas(currentCPU float64, currentReplicas int) int {
	cpuDifference := currentCPU - s.targetCPU
	adjustment := int(float64(currentReplicas) * cpuDifference)
	newReplicas := currentReplicas + adjustment
	if newReplicas < 1 {
		return 1
	}
	return newReplicas
}
