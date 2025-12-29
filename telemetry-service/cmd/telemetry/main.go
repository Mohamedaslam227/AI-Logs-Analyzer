package main
import (
	"log"
	"fmt"
	"telemetry-service/internal/config"
	"telemetry-service/internal/k8s"
	"telemetry-service/internal/scheduler"
)
func main() {
	fmt.Println("Main Function")
	cfg := config.Load()
	_ = cfg 
	client,err := k8s.NewClient()
	if err != nil {
		log.Fatal("Failed to create k8s client", err)
	}
	log.Println("Successfully created k8s client",client)
	scheduler := scheduler.New(cfg, client)
	scheduler.Start()
}