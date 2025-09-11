package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InterfaceExt struct {
	PreserveOnDelete *bool `json:"preserve_on_delete,omitempty"`

	PortState *string `json:"port_state,omitempty"`

	FixedIps *[]FixedIp `json:"fixed_ips,omitempty"`

	NetId *string `json:"net_id,omitempty"`

	PortId *string `json:"port_id,omitempty"`

	MacAddr *string `json:"mac_addr,omitempty"`

	DeleteOnTermination *bool `json:"delete_on_termination,omitempty"`

	DriverMode *string `json:"driver_mode,omitempty"`

	MinRate *int32 `json:"min_rate,omitempty"`

	MultiqueueNum *int32 `json:"multiqueue_num,omitempty"`

	PciAddress *string `json:"pci_address,omitempty"`
}

func (o InterfaceExt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InterfaceExt struct{}"
	}

	return strings.Join([]string{"InterfaceExt", string(data)}, " ")
}
