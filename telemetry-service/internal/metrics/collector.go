package metrics
import (
	"context"
	"time"
)

type MetricType string
const (
	MetricCPU MetricType ="cpu"
	MetricMemory MetricType ="memory"
	MetricNetwork MetricType ="network"
	MetricPod MetricType ="pod"
	MetricNode MetricType ="node"
)

type Metric struct {
	Type MetricType `json: "type"`
	Namespace string `json:"namespace,omitempty"`
	Resource string `json:"resource"`
	Value float64 `json:"value"`
	Timestamp time.Time `json:"timestamp"`
	Unit string `json:"unit"`
	Labels map[string]string `json:"labels,omitempty"`
}

type Collector interface {
	Name() string
	Collect(ctx context.Context) ([]Metric, error)
}