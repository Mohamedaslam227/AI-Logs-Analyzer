package metrics
import (
	"context"
	"time"
	"telemetry-service/internal/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CPUCollector struct {
	client *k8s.Client
}

func NewCPUCollector(client *k8s.Client) *CPUCollector {
	return &CPUCollector{
		client: client,
	}

}

func (c *CPUCollector) Name() string {
	return "cpu_collector"
}

func (c *CPUCollector) Collect(ctx context.Context) ([]Metric, error) {
	var metrics []Metric
	pods, err := c.client.Clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	now := time.Now()
	for _, pod := range pods.Items {
		if pod.Status.Phase != corev1.PodRunning {
			continue
		}
		totalCPU := resource.NewMilliQuantity(0, resource.DecimalSI)
		for _, container := range pod.Status.ContainerStatuses {
			if cpuReq, ok := container.Resources.Requests[corev1.ResourceCPU]; ok {
				totalCPU.Add(cpuReq)

				
			}
		}
		metrics = append(metrics, Metric{
			Type:      MetricCPU,
			Namespace: pod.Namespace,
			Resource:  pod.Name,
			Value:     float64(totalCPU.MilliValue()),
			Unit:      "millicores",
			Timestamp: now,
			Labels: map[string]string{
				"node": pod.Spec.NodeName,
			},
		})
	}
	return metrics, nil
}