package app

import "github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"

type Plugin interface {
	Init() error
	Finish()

	RegisterDevice(*read_parameter.IndividualOptions) error
	PushInstantIndicators(*read_parameter.IndividualOptions, *read_parameter.InstantIndicators) error
	RegisterInstantIndicators(*read_parameter.IndividualOptions, *read_parameter.InstantIndicators) error

	Name() string
}

var plugins []Plugin

func RegisterPlugin(plugin Plugin) {
	plugins = append(plugins, plugin)
}
