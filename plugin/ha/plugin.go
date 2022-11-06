//go:build ha

package ha

import (
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"
	"mercury-client/app"
	"mercury-client/plugin/mqtt"
)

type plugin struct {
	publisher *mqtt.Plugin
	prefix    string
}

func New(publisher *mqtt.Plugin, prefix string) app.Plugin {
	return &plugin{
		publisher: publisher,
		prefix:    prefix,
	}
}

func (r *plugin) Init() error {
	return nil
}

func (r *plugin) Finish() {

}

func (r *plugin) RegisterDevice(device *read_parameter.IndividualOptions) error {
	return nil
}

func (r *plugin) PushInstantIndicators(device *read_parameter.IndividualOptions, indicators *read_parameter.InstantIndicators) error {
	return nil
}

func (r *plugin) RegisterInstantIndicators(device *read_parameter.IndividualOptions, indicators *read_parameter.InstantIndicators) error {
	haDevice := NewDevice(*device)
	sensors := NewSensorInstantIndicators(
		haDevice,
		r.publisher.DeviceTopic(device)+mqtt.SuffixInstant,
		"value_json",
	)
	for _, s := range sensors {
		s.ExpireAfter = 300
		_ = r.publisher.Publish(r.prefix+s.Topic(), s, 1, true)
	}
	return nil
}

func (r *plugin) Name() string {
	return "HA plugin"
}
