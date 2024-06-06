package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeDetail 节点监控详情。
type NodeDetail struct {

	// 运行状态。
	RunningStatus *string `json:"running_status,omitempty"`

	// CPU使用率。
	CpuUsage *string `json:"cpu_usage,omitempty"`

	// 内存使用率。
	MemoryUsage *string `json:"memory_usage,omitempty"`

	// 硬盘使用率。
	DiskUsage *string `json:"disk_usage,omitempty"`

	// 总内存。单位MB。
	TotalMemory *string `json:"total_memory,omitempty"`

	// 可用内存。单位MB。
	AvailableMemory *string `json:"available_memory,omitempty"`

	// 总硬盘空间。单位GB。
	TotalHardDiskSpace *string `json:"total_hard_disk_space,omitempty"`

	// 可用硬盘空间。单位GB。
	AvailableHardDiskSpace *string `json:"available_hard_disk_space,omitempty"`

	// 网络读取速度。单位Byte/s。
	NetworkRead *string `json:"network_read,omitempty"`

	// 网络写入速度。单位Byte/s。
	NetworkWrite *string `json:"network_write,omitempty"`
}

func (o NodeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeDetail struct{}"
	}

	return strings.Join([]string{"NodeDetail", string(data)}, " ")
}
