package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Volume 逻辑盘的信息
type Volume struct {

	// 逻辑盘名称
	Name *string `json:"name,omitempty"`

	// RAID级别
	VolumeRaidLevel *string `json:"volume_raid_level,omitempty"`

	// 容量（单位：byte）
	CapacityBytes *int32 `json:"capacity_bytes,omitempty"`

	// 逻辑盘的条带大小（单位：byte）
	OptimumIoSizeBytes *int32 `json:"optimum_io_size_bytes,omitempty"`

	// 当前的读策略
	CurrentReadPolicy *string `json:"current_read_policy,omitempty"`

	// 默认的读策略
	DefaultReadPolicy *string `json:"default_read_policy,omitempty"`

	// 当前的写策略
	CurrentWritePolicy *string `json:"current_write_policy,omitempty"`

	// 默认的写策略
	DefaultWritePolicy *string `json:"default_write_policy,omitempty"`

	// 访问策略
	AccessPolicy *string `json:"access_policy,omitempty"`

	// 当前IO策略
	CurrentIoPolicy *string `json:"current_io_policy,omitempty"`

	// 当前IO策略
	DefaultIoPolicy *string `json:"default_io_policy,omitempty"`

	// 存储物理盘详细信息
	Drives *[]Drive `json:"drives,omitempty"`

	Status *Status `json:"status,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}
