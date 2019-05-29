package sensors

import (
	dht "github.com/d2r2/go-dht"
	"github.com/pkg/errors"
	"github.com/rikez/rpi-sensor-collector/infra/config"
)

// SensorMetrics originated by the DHT22
type SensorMetrics struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	DeviceID    string  `json:"deviceId"`
}

// DHTMetrics stands for the DHT sensor
type DHTMetrics struct{}

// Collect collects the metrics from the dht22 sensor
func (dhtImpl *DHTMetrics) Collect() (*SensorMetrics, error) {
	temperature, humidity, _, err := dht.ReadDHTxxWithRetry(dht.DHT22, 4, false, 10)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to collect metrics from the DHT22 sensors")
	}

	return &SensorMetrics{
		Temperature: temperature,
		Humidity:    humidity,
		DeviceID:    config.Envs.DeviceID,
	}, nil
}
