package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StorageController 存储控制器的信息
type StorageController struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 型号
	Model *string `json:"model,omitempty"`

	// 制造商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`

	// 健康情况
	Health *string `json:"health,omitempty"`

	// 存储控制器的类型
	Type *string `json:"type,omitempty"`

	// 固件版本
	FirmwareVersion *string `json:"firmware_version,omitempty"`

	// 支持的raid级别列表
	SupportedRaidLevels *[]string `json:"supported_raid_levels,omitempty"`

	// 存储控制器的模式
	Mode *string `json:"mode,omitempty"`

	// 储控制器配置版本
	ConfigurationVersion *string `json:"configuration_version,omitempty"`

	// SAS地址
	SasAddress *string `json:"sas_address,omitempty"`

	// 存储控制器BBU名称
	CapacitanceName *string `json:"capacitance_name,omitempty"`

	// 存储控制器电容(BBU)使能状态
	CapacitanceState *string `json:"capacitance_state,omitempty"`

	// 存储控制器电容(BBU)健康状态
	CapacitanceHealth *string `json:"capacitance_health,omitempty"`

	// 控制器支持最小条带值
	MinStripeSizeBytes *int64 `json:"min_stripe_size_bytes,omitempty"`

	// 控制器支持最大条带值
	MaxStripeSizeBytes *int64 `json:"max_stripe_size_bytes,omitempty"`

	// 逻辑盘列表
	Volumes *[]Volume `json:"volumes,omitempty"`

	// 管理的驱动器列表信息
	Drives *[]string `json:"drives,omitempty"`
}

func (o StorageController) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StorageController struct{}"
	}

	return strings.Join([]string{"StorageController", string(data)}, " ")
}
