package main

import (
	"os"
	"os/signal"

	"github.com/rikez/rpi-sensor-collector/sensors"

	"github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"github.com/rikez/rpi-sensor-collector/infra/config"
	"github.com/rikez/rpi-sensor-collector/infra/logger"
	"github.com/rikez/rpi-sensor-collector/metrics"
)

func init() {
	logger.SetupLogger()

	if err := config.ValidateEnv(); err != nil {
		err = errors.Wrap(err, "Failed to validate the environment variables")
		logger.Fatal(err)
	}
}

func main() {
	abortCh := make(chan struct{})

	// Starting the event stream
	notifyCh := make(chan metrics.Notification)
	metrics.StartEventCollector(&sensors.DHTMetrics{}, config.Envs.CollectorFreq, notifyCh)
	go func() {
		for {
			select {
			case ev := <-notifyCh:
				if ev.Error != nil {
					logger.Error(errors.Wrapf(ev.Error, "Received an error notification from the metrics collector: %s", ev.Message))
					continue
				}

				logger.WithFields(logrus.Fields{
					"event_message": ev.Message,
					"metrics":       *ev.Metrics,
				}).Info("Received metrics collector event")

				// Doing HTTP Request
				metrics.Send(ev.Metrics)

			case <-abortCh:
				return
			}
		}
	}()

	// Listening for OS signals for a graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// Handling the ERROR from the bowl and the QUIT channel from the OS
	logger.Info("Waiting for events or signals that can stop the application...")
	select {
	case <-quit:
		close(abortCh)
		logger.Info("Disposing the rpi-sensor-collector...")
	}

	logger.Info("Disposed the rpi-sensor-collector")
}
