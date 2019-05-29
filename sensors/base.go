package sensors

// Sensor common interface
type Sensor interface {
	Collect() (*SensorMetrics, error)
}
