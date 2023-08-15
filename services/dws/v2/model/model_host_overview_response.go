package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostOverviewResponse struct {

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 主机状态
	HostStat *string `json:"host_stat,omitempty"`

	// IP地址
	WorkIp *string `json:"work_ip,omitempty"`

	// 系统中未使用的内存(GB)。
	MemFree *float64 `json:"mem_free,omitempty"`

	// 总内存(GB)。
	MemTotal *float64 `json:"mem_total,omitempty"`

	// 内存使用率(GB)。
	MemUsage *float64 `json:"mem_usage,omitempty"`

	// 缓存内存(GB)。
	MemCached *float64 `json:"mem_cached,omitempty"`

	// 缓冲内存(MB)。
	MemBuffer *float64 `json:"mem_buffer,omitempty"`

	// ram暂存在swap中的大小(GB)。
	SwapFree *float64 `json:"swap_free,omitempty"`

	// 交换空间总和(GB)。
	SwapTotal *float64 `json:"swap_total,omitempty"`

	// CPU使用率(%)。
	CpuUsage *float64 `json:"cpu_usage,omitempty"`

	// 系统CPU占用率(%)。
	CpuUsageSys *float64 `json:"cpu_usage_sys,omitempty"`

	// 用户CPU占用率(%)。
	CpuUsageUsr *float64 `json:"cpu_usage_usr,omitempty"`

	// 空闲CPU占用率(%)。
	CpuIdle *float64 `json:"cpu_idle,omitempty"`

	// IO等待(%)。
	CpuIowait *float64 `json:"cpu_iowait,omitempty"`

	// 磁盘平均使用率(%)。
	DiskUsageAvg *float64 `json:"disk_usage_avg,omitempty"`

	// 磁盘总容量(GB)。
	DiskTotal *float64 `json:"disk_total,omitempty"`

	// 磁盘使用容量(GB)。
	DiskUsed *float64 `json:"disk_used,omitempty"`

	// 磁盘可用容量(GB)。
	DiskAvailable *float64 `json:"disk_available,omitempty"`

	// 磁盘IO(KB/s)。
	DiskIo *float64 `json:"disk_io,omitempty"`

	// 磁盘读速率(KB/s)。
	DiskIoRead *float64 `json:"disk_io_read,omitempty"`

	// 磁盘写速率(KB/s)。
	DiskIoWrite *float64 `json:"disk_io_write,omitempty"`

	// TCP协议栈重传率(%)。
	TcpResendRate *float64 `json:"tcp_resend_rate,omitempty"`

	// 网络IO(KB/s)。
	NetIo *float64 `json:"net_io,omitempty"`
}

func (o HostOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostOverviewResponse struct{}"
	}

	return strings.Join([]string{"HostOverviewResponse", string(data)}, " ")
}
