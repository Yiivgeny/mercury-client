//go:build mqtt || ha

package mqtt

import (
	"encoding/json"
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"mercury-client/app"
)

const PrefixDevice = "mercury/"
const SuffixInstant = "/instant"

type Plugin struct {
	logger *zap.Logger
	client mqtt.Client
}

func New(logger *zap.Logger, brokers []string) *Plugin {
	clientOptions := mqtt.
		NewClientOptions().
		SetClientID(app.Name)

	for _, broker := range brokers {
		clientOptions.AddBroker(broker)
	}
	client := mqtt.NewClient(clientOptions)

	p := &Plugin{
		logger: logger,
		client: client,
	}

	mqtt.ERROR = NewZapWrapper(logger, zap.ErrorLevel)
	mqtt.CRITICAL = NewZapWrapper(logger, zap.FatalLevel)
	mqtt.WARN = NewZapWrapper(logger, zap.WarnLevel)
	mqtt.DEBUG = NewZapWrapper(logger, zap.DebugLevel)

	return p
}

func (_ *Plugin) Name() string {
	return "MQTT Plugin"
}

func (r *Plugin) Init() error {
	if token := r.client.Connect(); token.Wait() && token.Error() != nil {
		return errors.WithStack(token.Error())
	}
	return nil
}

func (r *Plugin) Finish() {
	r.client.Disconnect(250)
}

func (r *Plugin) RegisterDevice(device *read_parameter.IndividualOptions) error {
	if err := r.Publish(r.DeviceTopic(device), device, 1, true); err != nil {
		return err
	}
	return nil
}

func (r *Plugin) PushInstantIndicators(device *read_parameter.IndividualOptions, indicators *read_parameter.InstantIndicators) error {
	if err := r.Publish(r.DeviceTopic(device)+SuffixInstant, indicators, 0, false); err != nil {
		return err
	}
	return nil
}

func (r *Plugin) RegisterInstantIndicators(device *read_parameter.IndividualOptions, indicators *read_parameter.InstantIndicators) error {
	return nil
}

func (r *Plugin) DeviceTopic(device *read_parameter.IndividualOptions) string {
	return PrefixDevice + device.SerialNumber
}

func (r *Plugin) Publish(topic string, v interface{}, qos byte, retained bool) error {
	text, err := json.Marshal(v)
	if err != nil {
		return err
	}
	token := r.client.Publish(topic, qos, retained, text)
	if token.Error() != nil {
		return errors.WithStack(token.Error())
	}
	token.Wait()
	if token.Error() != nil {
		return errors.WithStack(token.Error())
	}

	return nil
}
