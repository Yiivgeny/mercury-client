package ha

import (
	"fmt"
	"strings"
)

type Sensor struct {
	Device            Device `json:"device"`
	Name              string `json:"name"`
	UniqueId          string `json:"unique_id"`
	ObjectId          string `json:"object_id"`
	StateTopic        string `json:"state_topic"`
	DeviceClass       string `json:"device_class,omitempty"`
	StateClass        string `json:"state_class,omitempty"`
	UnitOfMeasurement string `json:"unit_of_measurement,omitempty"`
	ValueTemplate     string `json:"value_template"`
	ForceUpdate       bool   `json:"force_update,omitempty"`
	ExpireAfter       int    `json:"expire_after,omitempty"`

	topic string
}

func (s *Sensor) Topic() string {
	return s.topic
}

func NewSensorInstantIndicators(device Device, topic string, prefix string) []Sensor {
	sensor := make([]Sensor, 0)

	sensor = append(sensor, NewSensorVoltage(device, topic, prefix+".voltage")...)
	sensor = append(sensor, NewSensorCurrent(device, topic, prefix+".current")...)
	sensor = append(sensor, NewSensorPowerActive(device, topic, prefix+".power_p")...)
	sensor = append(sensor, NewSensorPowerReactive(device, topic, prefix+".power_q")...)
	sensor = append(sensor, NewSensorPowerFactor(device, topic, prefix+".power_factor")...)
	sensor = append(sensor, NewSensorTemperature(device, topic, prefix+".temperature"))
	sensor = append(sensor, NewSensorFrequency(device, topic, prefix+".frequency"))

	return sensor
}

func NewSensorEnergyPhase(device Device, topic string) []Sensor {
	uq := device.Identifiers[0]
	a := Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_energy_total_a", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_energy_total_a", uq),
		Name:              "Total Energy Phase A",
		StateTopic:        topic,
		DeviceClass:       "energy",
		StateClass:        "total_increasing",
		UnitOfMeasurement: "Wh",
		ValueTemplate:     "{{ value_json.a }}",
		ForceUpdate:       true,

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/energy_total_a/config", device.Manufacturer, uq)),
	}

	b := a
	b.UniqueId = fmt.Sprintf("mercury_%s_energy_total_b", uq)
	b.ObjectId = fmt.Sprintf("mercury_%s_energy_total_b", uq)
	b.Name = "Total Energy Phase B"
	b.ValueTemplate = "{{ value_json.b }}"
	b.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/energy_total_b/config", device.Manufacturer, uq))

	c := a
	c.UniqueId = fmt.Sprintf("mercury_%s_energy_total_c", uq)
	c.ObjectId = fmt.Sprintf("mercury_%s_energy_total_c", uq)
	c.Name = "Total Energy Phase C"
	c.ValueTemplate = "{{ value_json.c }}"
	c.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/energy_total_c/config", device.Manufacturer, uq))

	return []Sensor{a, b, c}
}

func NewSensorVoltage(device Device, topic string, prefix string) []Sensor {
	uq := device.Identifiers[0]

	a := Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_voltage_a", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_voltage_a", uq),
		Name:              "Voltage Phase A",
		StateTopic:        topic,
		DeviceClass:       "voltage",
		StateClass:        "measurement",
		UnitOfMeasurement: "V",
		ValueTemplate:     fmt.Sprintf("{{ %s.a }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/voltage_a/config", device.Manufacturer, uq)),
	}

	b := a
	b.UniqueId = fmt.Sprintf("mercury_%s_voltage_b", uq)
	b.ObjectId = fmt.Sprintf("mercury_%s_voltage_b", uq)
	b.Name = "Voltage Phase B"
	b.ValueTemplate = fmt.Sprintf("{{ %s.b }}", prefix)
	b.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/voltage_b/config", device.Manufacturer, uq))

	c := a
	c.UniqueId = fmt.Sprintf("mercury_%s_voltage_c", uq)
	c.ObjectId = fmt.Sprintf("mercury_%s_voltage_c", uq)
	c.Name = "Voltage Phase C"
	c.ValueTemplate = fmt.Sprintf("{{ %s.c }}", prefix)
	c.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/voltage_c/config", device.Manufacturer, uq))

	return []Sensor{a, b, c}
}

func NewSensorCurrent(device Device, topic string, prefix string) []Sensor {
	uq := device.Identifiers[0]

	a := Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_current_a", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_current_a", uq),
		Name:              "Current Phase A",
		StateTopic:        topic,
		DeviceClass:       "current",
		StateClass:        "measurement",
		UnitOfMeasurement: "A",
		ValueTemplate:     fmt.Sprintf("{{ %s.a }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/current_a/config", device.Manufacturer, uq)),
	}

	b := a
	b.UniqueId = fmt.Sprintf("mercury_%s_current_b", uq)
	b.ObjectId = fmt.Sprintf("mercury_%s_current_b", uq)
	b.Name = "Current Phase B"
	b.ValueTemplate = fmt.Sprintf("{{ %s.b }}", prefix)
	b.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/current_b/config", device.Manufacturer, uq))

	c := a
	c.UniqueId = fmt.Sprintf("mercury_%s_current_c", uq)
	c.ObjectId = fmt.Sprintf("mercury_%s_current_c", uq)
	c.Name = "Current Phase C"
	c.ValueTemplate = fmt.Sprintf("{{ %s.c }}", prefix)
	c.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/current_c/config", device.Manufacturer, uq))

	return []Sensor{a, b, c}
}

func NewSensorPowerActive(device Device, topic string, prefix string) []Sensor {
	uq := device.Identifiers[0]

	a := Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_power_p_a", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_power_p_a", uq),
		Name:              "Active power Phase A",
		StateTopic:        topic,
		DeviceClass:       "power",
		StateClass:        "measurement",
		UnitOfMeasurement: "W",
		ValueTemplate:     fmt.Sprintf("{{ %s.a }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_p_a/config", device.Manufacturer, uq)),
	}

	b := a
	b.UniqueId = fmt.Sprintf("mercury_%s_power_p_b", uq)
	b.ObjectId = fmt.Sprintf("mercury_%s_power_p_b", uq)
	b.Name = "Active power Phase B"
	b.ValueTemplate = fmt.Sprintf("{{ %s.b }}", prefix)
	b.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_p_b/config", device.Manufacturer, uq))

	c := a
	c.UniqueId = fmt.Sprintf("mercury_%s_power_p_c", uq)
	c.ObjectId = fmt.Sprintf("mercury_%s_power_p_c", uq)
	c.Name = "Active power Phase C"
	c.ValueTemplate = fmt.Sprintf("{{ %s.c }}", prefix)
	c.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_p_c/config", device.Manufacturer, uq))

	t := a
	t.UniqueId = fmt.Sprintf("mercury_%s_power_p", uq)
	t.ObjectId = fmt.Sprintf("mercury_%s_power_p", uq)
	t.Name = "Power active"
	t.ValueTemplate = fmt.Sprintf("{{ %s.sum }}", prefix)
	t.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_p/config", device.Manufacturer, uq))

	return []Sensor{a, b, c, t}
}

func NewSensorPowerReactive(device Device, topic string, prefix string) []Sensor {
	uq := device.Identifiers[0]

	a := Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_power_q_a", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_power_q_a", uq),
		Name:              "Reactive power Phase A",
		StateTopic:        topic,
		DeviceClass:       "reactive_power",
		StateClass:        "measurement",
		UnitOfMeasurement: "W",
		ValueTemplate:     fmt.Sprintf("{{ %s.a }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_q_a/config", device.Manufacturer, uq)),
	}

	b := a
	b.UniqueId = fmt.Sprintf("mercury_%s_power_q_b", uq)
	b.ObjectId = fmt.Sprintf("mercury_%s_power_q_b", uq)
	b.Name = "Reactive power Phase B"
	b.ValueTemplate = fmt.Sprintf("{{ %s.b }}", prefix)
	b.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_q_b/config", device.Manufacturer, uq))

	c := a
	c.UniqueId = fmt.Sprintf("mercury_%s_power_q_c", uq)
	c.ObjectId = fmt.Sprintf("mercury_%s_power_q_c", uq)
	c.Name = "Reactive power Phase C"
	c.ValueTemplate = fmt.Sprintf("{{ %s.c }}", prefix)
	c.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_q_c/config", device.Manufacturer, uq))

	t := a
	t.UniqueId = fmt.Sprintf("mercury_%s_power_q", uq)
	t.ObjectId = fmt.Sprintf("mercury_%s_power_q", uq)
	t.Name = "Reactive power"
	t.ValueTemplate = fmt.Sprintf("{{ %s.sum }}", prefix)
	t.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_q/config", device.Manufacturer, uq))

	return []Sensor{a, b, c, t}
}

func NewSensorPowerFactor(device Device, topic string, prefix string) []Sensor {
	uq := device.Identifiers[0]

	a := Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_power_factor_a", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_power_factor_a", uq),
		Name:              "Power factor Phase A",
		StateTopic:        topic,
		DeviceClass:       "power_factor",
		StateClass:        "measurement",
		UnitOfMeasurement: "%",
		ValueTemplate:     fmt.Sprintf("{{ (%s.a * 100) | round(1) }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_factor_a/config", device.Manufacturer, uq)),
	}

	b := a
	b.UniqueId = fmt.Sprintf("mercury_%s_power_factor_b", uq)
	b.ObjectId = fmt.Sprintf("mercury_%s_power_factor_b", uq)
	b.Name = "Power factor Phase B"
	b.ValueTemplate = fmt.Sprintf("{{ (%s.b * 100) | round(1) }}", prefix)
	b.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_factor_b/config", device.Manufacturer, uq))

	c := a
	c.UniqueId = fmt.Sprintf("mercury_%s_power_factor_c", uq)
	c.ObjectId = fmt.Sprintf("mercury_%s_power_factor_c", uq)
	c.Name = "Power factor Phase C"
	c.ValueTemplate = fmt.Sprintf("{{ (%s.c * 100) | round(1) }}", prefix)
	c.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_factor_c/config", device.Manufacturer, uq))

	t := a
	t.UniqueId = fmt.Sprintf("mercury_%s_power_factor", uq)
	t.ObjectId = fmt.Sprintf("mercury_%s_power_factor", uq)
	t.Name = "Power factor"
	t.ValueTemplate = fmt.Sprintf("{{ (%s.sum * 100) | round(1) }}", prefix)
	t.topic = strings.ToLower(fmt.Sprintf("/sensor/%s_%s/power_factor/config", device.Manufacturer, uq))

	return []Sensor{a, b, c, t}
}

func NewSensorTemperature(device Device, topic string, prefix string) Sensor {
	uq := device.Identifiers[0]

	return Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_temperature", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_temperature", uq),
		Name:              "Temperature",
		StateTopic:        topic,
		DeviceClass:       "temperature",
		StateClass:        "measurement",
		UnitOfMeasurement: "°C",
		ValueTemplate:     fmt.Sprintf("{{ %s.temperature }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/temperature/config", device.Manufacturer, uq)),
	}

}

func NewSensorFrequency(device Device, topic string, prefix string) Sensor {
	uq := device.Identifiers[0]

	return Sensor{
		Device:            device,
		UniqueId:          fmt.Sprintf("mercury_%s_frequency", uq),
		ObjectId:          fmt.Sprintf("mercury_%s_frequency", uq),
		Name:              "Frequency",
		StateTopic:        topic,
		DeviceClass:       "frequency",
		StateClass:        "measurement",
		UnitOfMeasurement: "Hz",
		ValueTemplate:     fmt.Sprintf("{{ %s.frequency }}", prefix),

		topic: strings.ToLower(fmt.Sprintf("/sensor/%s_%s/frequency/config", device.Manufacturer, uq)),
	}

}
