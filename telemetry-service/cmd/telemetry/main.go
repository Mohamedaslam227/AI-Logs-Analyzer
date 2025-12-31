package main
import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"telemetry-service/internal/config"
	"telemetry-service/internal/k8s"
	"telemetry-service/internal/scheduler"
)
func main() {
	cfg := config.Load()
	_ = cfg 
	client,err := k8s.NewClient()
	if err != nil {
		log.Fatal("Failed to create k8s client", err)
	}
	log.Println("Successfully created k8s client",client)
	scheduler := scheduler.New(cfg, client)
	scheduler.Start()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	log.Printf("Received signal: %s", sig.String())
	log.Println("Initiating graceful shutdown...")
	scheduler.Stop()
	time.Sleep(500 * time.Millisecond)


}