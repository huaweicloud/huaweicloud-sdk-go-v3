package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DiskInfo 磁盘使用情况
type DiskInfo struct {

	// 数据盘可用量
	DataAvailable *float64 `json:"data_available,omitempty"`

	// 数据盘总容量
	DataTotal *float64 `json:"data_total,omitempty"`

	// 数据盘使用百分比
	DataUsePercent *int32 `json:"data_use_percent,omitempty"`

	// 记录ID
	Id *string `json:"id,omitempty"`

	// 临时目录可用量
	SwapAvailable *float64 `json:"swap_available,omitempty"`

	// 临时目录总容量
	SwapTotal *float64 `json:"swap_total,omitempty"`

	// 临时目录使用百分比
	SwapUsePercent *int32 `json:"swap_use_percent,omitempty"`

	// 系统盘可用量
	SystemAvailable *float64 `json:"system_available,omitempty"`

	// 系统盘容量
	SystemTotal *float64 `json:"system_total,omitempty"`

	// 系统盘使用百分比
	SystemUsePercent *int32 `json:"system_use_percent,omitempty"`

	// 记录时间
	Time *string `json:"time,omitempty"`
}

func (o DiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskInfo struct{}"
	}

	return strings.Join([]string{"DiskInfo", string(data)}, " ")
}
