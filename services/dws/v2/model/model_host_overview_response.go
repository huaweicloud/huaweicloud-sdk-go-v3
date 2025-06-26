package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HostOverviewResponse struct {

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 主机名称。 **取值范围**： 不涉及。
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 主机状态。 **取值范围**： 不涉及。
	HostStat *string `json:"host_stat,omitempty"`

	// **参数解释**： IP地址。 **取值范围**： 不涉及。
	WorkIp *string `json:"work_ip,omitempty"`

	// **参数解释**： 系统中未使用的内存(GB)。 **取值范围**： 不涉及。
	MemFree *float64 `json:"mem_free,omitempty"`

	// **参数解释**： 总内存(GB)。 **取值范围**： 不涉及。
	MemTotal *float64 `json:"mem_total,omitempty"`

	// **参数解释**： 内存使用率(GB)。 **取值范围**： 不涉及。
	MemUsage *float64 `json:"mem_usage,omitempty"`

	// **参数解释**： 缓存内存(GB)。 **取值范围**： 不涉及。
	MemCached *float64 `json:"mem_cached,omitempty"`

	// **参数解释**： 缓冲内存(MB)。 **取值范围**： 不涉及。
	MemBuffer *float64 `json:"mem_buffer,omitempty"`

	// **参数解释**： ram暂存在swap中的大小(GB)。 **取值范围**： 不涉及。
	SwapFree *float64 `json:"swap_free,omitempty"`

	// **参数解释**： 交换空间总和(GB)。 **取值范围**： 不涉及。
	SwapTotal *float64 `json:"swap_total,omitempty"`

	// **参数解释**： CPU使用率(%)。 **取值范围**： 不涉及。
	CpuUsage *float64 `json:"cpu_usage,omitempty"`

	// **参数解释**： 系统CPU占用率(%)。 **取值范围**： 不涉及。
	CpuUsageSys *float64 `json:"cpu_usage_sys,omitempty"`

	// **参数解释**： 用户CPU占用率(%)。 **取值范围**： 不涉及。
	CpuUsageUsr *float64 `json:"cpu_usage_usr,omitempty"`

	// **参数解释**： 空闲CPU占用率(%)。 **取值范围**： 不涉及。
	CpuIdle *float64 `json:"cpu_idle,omitempty"`

	// **参数解释**： IO等待(%)。 **取值范围**： 不涉及。
	CpuIowait *float64 `json:"cpu_iowait,omitempty"`

	// **参数解释**： 磁盘平均使用率(%)。 **取值范围**： 不涉及。
	DiskUsageAvg *float64 `json:"disk_usage_avg,omitempty"`

	// **参数解释**： 磁盘总容量(GB)。 **取值范围**： 不涉及。
	DiskTotal *float64 `json:"disk_total,omitempty"`

	// **参数解释**： 磁盘使用容量(GB)。 **取值范围**： 不涉及。
	DiskUsed *float64 `json:"disk_used,omitempty"`

	// **参数解释**： 磁盘可用容量(GB)。 **取值范围**： 不涉及。
	DiskAvailable *float64 `json:"disk_available,omitempty"`

	// **参数解释**： 磁盘IO(KB/s)。 **取值范围**： 不涉及。
	DiskIo *float64 `json:"disk_io,omitempty"`

	// **参数解释**： 磁盘读速率(KB/s)。 **取值范围**： 不涉及。
	DiskIoRead *float64 `json:"disk_io_read,omitempty"`

	// **参数解释**： 磁盘写速率(KB/s)。 **取值范围**： 不涉及。
	DiskIoWrite *float64 `json:"disk_io_write,omitempty"`

	// **参数解释**： TCP协议栈重传率(%)。 **取值范围**： 不涉及。
	TcpResendRate *float64 `json:"tcp_resend_rate,omitempty"`

	// **参数解释**： 网络IO(KB/s)。 **取值范围**： 不涉及。
	NetIo *float64 `json:"net_io,omitempty"`
}

func (o HostOverviewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostOverviewResponse struct{}"
	}

	return strings.Join([]string{"HostOverviewResponse", string(data)}, " ")
}
