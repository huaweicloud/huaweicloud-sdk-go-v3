package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Processors 处理器的信息，包含处理器的型号、核心数、线程数等。
type Processors struct {

	// 处理器名称
	Name *string `json:"name,omitempty"`

	// 制造商
	Manufacturer *string `json:"manufacturer,omitempty"`

	// 处理器型号
	Model *string `json:"model,omitempty"`

	// 处理器的总核数
	TotalCores *int32 `json:"total_cores,omitempty"`

	// 启用的核心数
	TotalEnabledCores *int32 `json:"total_enabled_cores,omitempty"`

	// 处理器的总线程数
	TotalThreads *int32 `json:"total_threads,omitempty"`

	// 启用的总线程
	TotalEnabledThreads *int32 `json:"total_enabled_threads,omitempty"`

	// 处理器的插槽号
	Socket *int32 `json:"socket,omitempty"`

	// 处理器的最大主频（单位：MHz）
	MaxSpeedMhz *int32 `json:"max_speed_mhz,omitempty"`

	// 处理器的温度
	Temperature *int32 `json:"temperature,omitempty"`

	// 一级缓存（单位：KiB）
	L1CacheKib *int32 `json:"l1_cache_kib,omitempty"`

	// 二级缓存（单位：KiB）
	L2CacheKib *int32 `json:"l2_cache_kib,omitempty"`

	// 三级缓存（单位：KiB）
	L3CacheKib *int32 `json:"l3_cache_kib,omitempty"`

	// 处理器的主频
	FrequencyMhz *int32 `json:"frequency_mhz,omitempty"`

	// 其他参数
	OtherParameters *string `json:"other_parameters,omitempty"`

	// 序列号
	SerialNumber *string `json:"serial_number,omitempty"`

	// 健康状态
	Health *string `json:"health,omitempty"`

	// 状态
	State *string `json:"state,omitempty"`
}

func (o Processors) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Processors struct{}"
	}

	return strings.Join([]string{"Processors", string(data)}, " ")
}
