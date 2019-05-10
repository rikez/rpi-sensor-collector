package main

import (
	"github.com/pkg/errors"
	"github.com/rikez/rpi-sensor-collector/infra/config"
	"github.com/rikez/rpi-sensor-collector/infra/logger"
	"github.com/sirupsen/logrus"
)

func init() {
	logger.SetupLogger()

	if err := config.ValidateEnv(); err != nil {
		err = errors.Wrap(err, "Failed to validate the environment variables")
		logrus.Fatal(err)
	}
}

func main() {

}
