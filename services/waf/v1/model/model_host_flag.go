package model

import (
	"encoding/json"

	"strings"
)

type HostFlag struct {
	// true/false

	PciDss *string `json:"pci_dss,omitempty"`
	// true/false

	Pci3ds *string `json:"pci_3ds,omitempty"`
	// old/new

	Cname *string `json:"cname,omitempty"`
	// true/false

	IsDualAz *string `json:"is_dual_az,omitempty"`
	// true/false

	Ipv6 *string `json:"ipv6,omitempty"`
}

func (o HostFlag) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HostFlag struct{}"
	}

	return strings.Join([]string{"HostFlag", string(data)}, " ")
}
