package main

import (
	"log"
	"os"
	"strconv"

	"github.com/egocentri/go-dispcalc1/internal/agent"
	"github.com/egocentri/go-dispcalc1/internal/config"
)

func main() {
	cfg := config.InitEnv()
	computingPower := 2
	if val := os.Getenv("COMPUTING_POWER"); val != "" {
		if cp, err := strconv.Atoi(val); err == nil {
			computingPower = cp
		}
	}
	orchestratorURL := "http://localhost:8080"
	if val := os.Getenv("ORCHESTRATOR_URL"); val != "" {
		orchestratorURL = val
	}

	log.Printf("Agent started. Orchestrator URL: %s, computingPower: %d\n", orchestratorURL, computingPower)
	worker := agent.NewWorker(orchestratorURL, computingPower, cfg)
	for {
		worker.RequestAndProcessTask()
	}
}
