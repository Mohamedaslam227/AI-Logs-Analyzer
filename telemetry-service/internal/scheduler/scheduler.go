package scheduler

import (
	"context"
	"time"
	"log"
	"sync"
	"telemetry-service/internal/config"
	"telemetry-service/internal/k8s"
)

type Scheduler struct {
	cfg *config.Config
	client *k8s.Client
	ctx context.Context
	cancel context.CancelFunc
	wg sync.WaitGroup
}


func New(cfg *config.Config, client *k8s.Client) *Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &Scheduler{
		cfg: cfg,
		client: client,
		ctx: ctx,
		cancel: cancel,
	}

}

func (s *Scheduler) Start() {
	log.Println("Starting Scheduler....!")
	s.wg.Add(1)
	go s.run()
}

func (s *Scheduler) Stop() {
	log.Println("Stopping Scheduler....!")
	s.cancel()
	s.wg.Wait()
	log.Println("Scheduler stopped.")
}

func (s *Scheduler) run() {
	defer s.wg.Done()
	ticker := time.NewTicker(s.cfg.PollInterval)
	defer ticker.Stop()
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.executeCycle()
		}
	}
}

func (s *Scheduler) executeCycle() {
	start := time.Now()
	log.Println("Executing cycle at", start)
	// Placeholder: collectors & detectors will be wired here
	// Example:
	// metrics := collectMetrics(s.client)
	// signals := detectIncidents(metrics)
	// publishEvents(signals)
	elapsed := time.Since(start)
	log.Println("Cycle completed in", elapsed)

}