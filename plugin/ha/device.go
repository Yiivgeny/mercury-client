package ha

import (
	"github.com/Yiivgeny/incotex-mercury-client/client/methods/read_parameter"
)

type Device struct {
	Name            string   `json:"name,omitempty"`
	Manufacturer    string   `json:"manufacturer,omitempty"`
	Model           string   `json:",omitempty"`
	Identifiers     []string `json:"identifiers,omitempty"`
	SoftwareVersion string   `json:"sw_version,omitempty"`
}

func NewDevice(individual read_parameter.IndividualOptions) Device {
	return Device{
		Name:         "Mercury №" + individual.SerialNumber,
		Manufacturer: "Mercury",
		Identifiers: []string{
			individual.SerialNumber,
		},
		SoftwareVersion: individual.Firmware,
	}
}
