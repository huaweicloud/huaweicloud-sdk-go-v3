package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NetworkAdapter 网络适配器的详细信息，如型号、厂商等
type NetworkAdapter struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 网络适配器的芯片型号
	CardModel *string `json:"card_model,omitempty"`

	// 网络适配器的型号
	Model *string `json:"model,omitempty"`

	// 网络适配器的制造商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 网络适配器的芯片制造商
	CardManufacturer *string `json:"card_manufacturer,omitempty"`

	// 网络适配器的位置
	Position *string `json:"position,omitempty"`

	// 网络适配器的卡槽号
	SlotNumber *int32 `json:"slot_number,omitempty"`

	// 网络适配器的PCB版本
	PcbVersion *string `json:"pcb_version,omitempty"`

	// 网络适配器的驱动名称
	DriverName *string `json:"driver_name,omitempty"`

	// 网络适配器的驱动版本
	DriverVersion *string `json:"driver_version,omitempty"`

	// 网络适配器的资源归属
	AssociatedResource *string `json:"associated_resource,omitempty"`

	// 网络适配器的固件版本
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	// 健康状态
	Health *string `json:"health,omitempty"`

	// 网络端口列表
	NetworkPorts *[]NetworkPort `json:"network_ports,omitempty"`
}

func (o NetworkAdapter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetworkAdapter struct{}"
	}

	return strings.Join([]string{"NetworkAdapter", string(data)}, " ")
}
