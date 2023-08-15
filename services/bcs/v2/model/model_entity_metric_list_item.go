package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EntityMetricListItem 监控数据列表项目
type EntityMetricListItem struct {

	// cpu使用率
	CpuUsage *string `json:"cpuUsage,omitempty"`

	// 磁盘读取速率
	DiskReadRate *string `json:"diskReadRate,omitempty"`

	// 磁盘写入速率
	DiskWriteRate *string `json:"diskWriteRate,omitempty"`

	// 物理内存使用率
	MemUsage *string `json:"memUsage,omitempty"`

	// 下行BPs
	RecvBytesRate *string `json:"recvBytesRate,omitempty"`

	// 上行BPs
	SendBytesRate *string `json:"sendBytesRate,omitempty"`

	// 文件系统使用率
	FilesystemUsage *string `json:"filesystemUsage,omitempty"`
}

func (o EntityMetricListItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EntityMetricListItem struct{}"
	}

	return strings.Join([]string{"EntityMetricListItem", string(data)}, " ")
}
