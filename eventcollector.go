package main

import (
	dht "github.com/d2r2/go-dht"
	"github.com/rikez/rpi-sensor-collector/infra/logger"
	"github.com/rikez/rpi-sensor-collector/infra/config"
	"github.com/sirupsen/logrus"
)

type Notify {
	Message string
	Args interface{}
	Error error
}

type SensorMetrics {
	Temperature float32 `json:"temperature"`
	Humidity float32 `json:"humidity"`
	DeviceID string `json:"deviceId"`
}

func collectMetrics() (*SensorMetrics, error) {
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to collect metrics from the DHT22 sensors")
	}

	return &SensorMetrics{
		Temperature: temperature,
		Humidity: humidity,
		DeviceID: config.EnvConfig.DeviceID,
	}
}

// StartEventStream start streaming the metrics to a Kafka topic.
func StartEventStream(interval int, notifyCh chan<- *Notify) {
	go func() {
		for {
			metrics, err := collectMetrics()
			if err != nil {
				notifyCh <- &Notify{
					Message: "Failed to collect the sensor metrics",
					Error: err,
				}
			} else {
				// produce to topic

				notifyCh <- &Notify{
					Message: "Successfully sent the sensor metrics",
					Args: metrics
					Error: nil,
				}
			}
		}
	}
}