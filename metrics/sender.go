package metrics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rikez/rpi-sensor-collector/sensors"

	"github.com/pkg/errors"
	"github.com/rikez/rpi-sensor-collector/infra/config"
	"github.com/rikez/rpi-sensor-collector/infra/logger"
)

// Send sends the sensor metrics through HTTP Request
func Send(m *sensors.SensorMetrics) {
	b, err := json.Marshal(m)
	if err != nil {
		logger.Error(errors.Wrapf(err, "Failed to marshalize the metrics: %v", m))
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", config.Envs.APIURL, "dht22"), bytes.NewReader(b))
	req.Header.Add("Authorization", config.Envs.Token)
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}
	_, err = client.Do(req)
	if err != nil {
		logger.Error(errors.Wrap(err, "Failed to perform the HTTP request to the server"))
		return
	}

	logger.Infof("Request succeeded: %s", m.DeviceID)
}
