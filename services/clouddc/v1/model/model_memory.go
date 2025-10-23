package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Memory 内存条信息
type Memory struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 表示系统的总内存容量（单位：MiB）
	CapacityMib *int32 `json:"capacity_mib,omitempty"`

	// 制造商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 内存类型：DDR4/DDR6
	MemoryDeviceType *string `json:"memory_device_type,omitempty"`

	// 主频（单位：MHz）
	AllowedSpeedMhz *int32 `json:"allowed_speed_mhz,omitempty"`

	// 当前频率（单位：MHz）
	OperatingSpeedMhz *int32 `json:"operating_speed_mhz,omitempty"`

	// 序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 健康状态
	Health *string `json:"health,omitempty"`

	// 启用状态
	State *string `json:"state,omitempty"`

	// Rank数量
	RankCount *int32 `json:"rank_count,omitempty"`

	// 数据带宽
	DataWidthBits *int32 `json:"data_width_bits,omitempty"`

	// 部件编号
	PartNumber *string `json:"part_number,omitempty"`

	// 最小电压
	MinVoltageMillivolt *int32 `json:"min_voltage_millivolt,omitempty"`

	// Bom编码
	BomNumber *string `json:"bom_number,omitempty"`

	// 类型详细信息
	TypeDetail *string `json:"type_detail,omitempty"`

	// 技术
	Technology *string `json:"technology,omitempty"`

	// 位置
	Position *string `json:"position,omitempty"`
}

func (o Memory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Memory struct{}"
	}

	return strings.Join([]string{"Memory", string(data)}, " ")
}
