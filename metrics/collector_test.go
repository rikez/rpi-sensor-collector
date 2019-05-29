package metrics

import (
	"errors"
	"testing"
	"time"

	"github.com/rikez/rpi-sensor-collector/sensors"
)

//Mocking the sensor
type SensorMock struct{}

// Collect method mocked
func (dhtImpl *SensorMock) Collect() (*sensors.SensorMetrics, error) {

	return &sensors.SensorMetrics{
		Temperature: 10,
		Humidity:    10,
		DeviceID:    "12345",
	}, nil
}

func TestCollectorNotificationReturnSensorMetrics(t *testing.T) {
	abortCh := make(chan struct{})
	notifyCh := make(chan Notification)
	mock := SensorMock{}
	StartEventCollector(&mock, time.Duration(50), notifyCh)

	for {
		select {
		case n := <-notifyCh:
			if n.Metrics.Temperature != 10 {
				t.Fatalf("Temperature is incorrect, got: %f, want: %f.", n.Metrics.Temperature, 10.0)
			}
			if n.Metrics.Humidity != 10 {
				t.Fatalf("Humidity is incorrect, got: %f, want: %f.", n.Metrics.Humidity, 10.0)
			}
			if n.Metrics.DeviceID != "12345" {
				t.Fatalf("DeviceID is incorrect, got: %s, want: %s.", n.Metrics.DeviceID, "12345")
			}

			return
		case <-abortCh:
			return
		}
	}
}

//Mocking the sensor
type SensorMock2 struct{}

// Collect method mocked
func (dhtImpl *SensorMock2) Collect() (*sensors.SensorMetrics, error) {
	return nil, errors.New("Failed to collect the metrics")
}

func TestCollectorNotificationReturnError(t *testing.T) {
	abortCh := make(chan struct{})
	notifyCh := make(chan Notification)
	mock := SensorMock2{}
	StartEventCollector(&mock, time.Duration(50), notifyCh)

	for {
		select {
		case n := <-notifyCh:
			if n.Metrics != nil {
				t.Fatalf("Metrics are incorrect, got: %v, want: nil", n.Metrics)
			}
			if n.Error == nil {
				t.Fatalf("Error is incorrect, got: nil, want: %v", n.Error)
			}
			if n.Message != "Failed to collect the sensor metrics" {
				t.Fatalf("Message is incorrect, got: %s, want: Failed to collect the sensor metrics", n.Message)
			}

			return
		case <-abortCh:
			return
		}
	}
}
