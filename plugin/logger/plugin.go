package logger

import (
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"
	"go.uber.org/zap"
	"mercury-client/app"
)

type plugin struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) app.Plugin {
	return &plugin{
		logger: logger,
	}
}

func (r *plugin) Init() error {
	return nil
}

func (r *plugin) Finish() {

}

func (r *plugin) RegisterDevice(device *read_parameter.IndividualOptions) error {
	r.logger.Info(
		"Device detected",
		zap.String("number", device.SerialNumber),
		zap.Any("data", device),
	)
	return nil
}

func (r *plugin) PushInstantIndicators(device *read_parameter.IndividualOptions, indicators *read_parameter.InstantIndicators) error {
	r.logger.Info(
		"Instant parameters retrieved",
		zap.String("number", device.SerialNumber),
		zap.Any("indicators", indicators),
	)
	return nil
}

func (r *plugin) RegisterInstantIndicators(device *read_parameter.IndividualOptions, indicators *read_parameter.InstantIndicators) error {
	return nil
}

func (r *plugin) Name() string {
	return "Logger plugin"
}
