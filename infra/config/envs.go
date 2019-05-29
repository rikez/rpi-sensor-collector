package config

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// Stands for the used environment variables
type Env struct {
	Token         string
	APIURL        string
	DeviceID      string
	CollectorFreq time.Duration
}

const (
	APIURLConst   = "API_URL"
	Token         = "TOKEN"
	DeviceID      = "DEVICE_ID"
	CollectorFreq = "COLLECTOR_FREQUENCY"
)

// in-memory env variables
var Envs = &Env{}

// ValidateEnv check the required environment variables
func ValidateEnv() error {
	APIURL := os.Getenv(APIURLConst)
	if APIURL == "" {
		return errors.New("The APIURL variable is required")
	}
	deviceID := os.Getenv(DeviceID)
	if deviceID == "" {
		return errors.New("The DEVICE_ID variable is required")
	}
	token := os.Getenv(Token)
	if deviceID == "" {
		return errors.New("The TOKEN variable is required")
	}
	collectorFreqStr := os.Getenv(CollectorFreq)
	if collectorFreqStr == "" {
		return errors.New("The COLLECTOR_FREQUENCY variable is required")
	}
	collectorFreq, err := strconv.Atoi(collectorFreqStr)
	if err != nil {
		return errors.Wrap(err, "Failed to convert string to int")
	}
	if collectorFreq < 10 && collectorFreq > 10000 {
		return errors.New("The COLLECTOR_FREQUENCY variable is must be within 10 and 10000 Milliseconds")
	}

	Envs.Token = token
	Envs.APIURL = APIURL
	Envs.DeviceID = deviceID
	Envs.CollectorFreq = time.Duration(collectorFreq)

	return nil
}
