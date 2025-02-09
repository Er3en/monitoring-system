package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

var (
	// CPU metrics
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "system_cpu_usage",
			Help: "CPU usage percentage",
		},
		[]string{"cpu"},
	)
	// Memory metric
	memUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "system_memory_usage",
			Help: "Memory usage percentage",
		},
	)
)

func init() {
	log.Println("Initializing metrics...")
	metrics := []prometheus.Collector{cpuUsage, memUsage}
	for _, metric := range metrics {
		err := prometheus.Register(metric)
		if err != nil {
			log.Fatalf("Failed to register metric %v: %v", metric, err)
		}
	}
	log.Println("Metrics registered.")
}

// Function to collect metrics
func collectMetrics() {
	log.Println("Starting metric collection...")

	// Collect CPU metrics
	cpus, err := cpu.Percent(0, true)
	if err != nil {
		log.Println("Error collecting CPU metrics:", err)
		return
	}
	for i, cpuPercent := range cpus {
		log.Printf("Collected CPU metric for cpu%d: %.2f%%", i, cpuPercent)
		cpuUsage.WithLabelValues(fmt.Sprintf("cpu%d", i)).Set(cpuPercent)
	}
	totalCPUUsage, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("Error collecting CPU metrics:", err)
		return
	}
	log.Printf("Total CPU usage: %.2f%%", totalCPUUsage)

	// Collect memory metrics
	memStats, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error collecting memory metrics:", err)
		return
	}
	log.Printf("Collected memory metric: %.2f%%", memStats.UsedPercent)
	memUsage.Set(memStats.UsedPercent)

	log.Println("Metric collection completed.")
}

func main() {
	// Start collecting metrics in the background
	go func() {
		for {
			collectMetrics() // Collect metrics in the background
			log.Println("Waiting 5 seconds before the next metric collection...")
			time.Sleep(5 * time.Second)
		}
	}()

	// Endpoint to expose Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Endpoint /metrics started on port 8080")

	// Additional endpoint /collect to manually trigger metric collection
	http.HandleFunc("/collect", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Manual metric collection triggered...")
		collectMetrics()
		fmt.Fprintf(w, "Metrics have been collected.")
	})

	// Start HTTP server
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
