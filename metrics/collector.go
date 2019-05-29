package metrics

import (
	"time"

	"github.com/rikez/rpi-sensor-collector/sensors"
)

// Notification sent through the channel
type Notification struct {
	Message string
	Metrics *sensors.SensorMetrics
	Error   error
}

// StartEventCollector starts collecting the metrics.
func StartEventCollector(dht sensors.Sensor, interval time.Duration, notifyCh chan<- Notification) {
	go func() {
		for {
			time.Sleep(time.Millisecond * interval)

			metrics, err := dht.Collect()
			if err != nil {
				notifyCh <- Notification{
					Message: "Failed to collect the sensor metrics",
					Metrics: nil,
					Error:   err,
				}
				continue
			}

			// produce to topic
			notifyCh <- Notification{
				Message: "Produced the sensor metrics",
				Metrics: metrics,
				Error:   nil,
			}
		}
	}()
}
