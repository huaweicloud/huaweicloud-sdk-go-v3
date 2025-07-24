package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Power 电源详细信息
type Power struct {

	// 电源名称
	Name *string `json:"name,omitempty"`

	// 槽位
	SlotNumber *int32 `json:"slot_number,omitempty"`

	// 型号
	Model *string `json:"model,omitempty"`

	// 厂商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 输入模式
	PowerSupplyType *string `json:"power_supply_type,omitempty"`

	// 额定功率
	PowerCapacityWatts *string `json:"power_capacity_watts,omitempty"`

	// 输入电压
	LineInputVoltage *int32 `json:"line_input_voltage,omitempty"`

	// 输出电压
	OutputVoltage *int32 `json:"output_voltage,omitempty"`

	// 协议
	Protocol *string `json:"protocol,omitempty"`

	// 主备模式
	ActiveStandby *string `json:"active_standby,omitempty"`

	// 部件编号
	PartNumber *string `json:"part_number,omitempty"`

	// 序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 固件版本
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func (o Power) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Power struct{}"
	}

	return strings.Join([]string{"Power", string(data)}, " ")
}
