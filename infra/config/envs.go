package config

import (
	"os"

	"github.com/pkg/errors"
)

// Stands for the used environment variables
type Env struct {
	KafkaServers       string
	SensorMetricsTopic string
	DeviceID           string
}

const (
	KafkaServers       = "KAFKA_SERVERS"
	SensorMetricsTopic = "SENSOR_METRICS_TOPIC"
	DeviceID           = "DEVICE_ID"
)

// in-memory env variables
var EnvConfig = &Env{}

// ValidateEnv check the required environment variables
func ValidateEnv() error {
	kafkaServers := os.Getenv(KafkaServers)
	if kafkaServers == "" {
		return errors.New("The KAFKA_SERVERS variable is required")
	}

	sensorMetricsTopic := os.Getenv(SensorMetricsTopic)
	if kafkaServers == "" {
		return errors.New("The SENSOR_METRICS_TOPIC variable is required")
	}

	deviceID := os.Getenv(DeviceID)
	if deviceID == "" {
		return errors.New("The DEVICE_ID variable is required")
	}

	EnvConfig.KafkaServers = kafkaServers
	EnvConfig.SensorMetricsTopic = sensorMetricsTopic
	EnvConfig.DeviceID = deviceID

	return nil
}
