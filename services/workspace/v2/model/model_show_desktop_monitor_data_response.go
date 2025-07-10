package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopMonitorDataResponse Response Object
type ShowDesktopMonitorDataResponse struct {

	// 在线信息。
	UserOnlineInfo *[]MonitorUserOnlineInfo `json:"user_online_info,omitempty"`

	// 计算机登录状态，在线信息没值的情况下，查看计算机登录状态 0 未登录， 1 登录中, 2登录状态同当前。
	OnlineStatus *int32 `json:"online_status,omitempty"`

	// cpu使用率信息。
	CpuInfo *[]Datapoints `json:"cpu_info,omitempty"`

	// 内存使用率信息。
	MemoryInfo *[]Datapoints `json:"memory_info,omitempty"`

	// 磁盘使用率。
	DiskUtilInband *[]Datapoints `json:"disk_util_inband,omitempty"`

	// 磁盘读带宽。
	DiskReadBytesRate *[]Datapoints `json:"disk_read_bytes_rate,omitempty"`

	// 磁盘写带宽。
	DiskWriteBytesRate *[]Datapoints `json:"disk_write_bytes_rate,omitempty"`

	// 磁盘读IOPS。
	DiskReadRequestsRate *[]Datapoints `json:"disk_read_requests_rate,omitempty"`

	// 磁盘写IOPS。
	DiskWriteRequestsRate *[]Datapoints `json:"disk_write_requests_rate,omitempty"`

	// 带内网络流入速率。
	NetworkIncomingBytesRateInband *[]Datapoints `json:"network_incoming_bytes_rate_inband,omitempty"`

	// 带内网络流出速率。
	NetworkOutgoingBytesRateInband *[]Datapoints `json:"network_outgoing_bytes_rate_inband,omitempty"`

	// 带外网络流入速率。
	NetworkIncomingBytesAggregateRate *[]Datapoints `json:"network_incoming_bytes_aggregate_rate,omitempty"`

	// 带外网络流出速率。
	NetworkOutgoingBytesAggregateRate *[]Datapoints `json:"network_outgoing_bytes_aggregate_rate,omitempty"`

	// 网络连接数。
	NetworkVmConnections *[]Datapoints `json:"network_vm_connections,omitempty"`
	HttpStatusCode       int           `json:"-"`
}

func (o ShowDesktopMonitorDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopMonitorDataResponse struct{}"
	}

	return strings.Join([]string{"ShowDesktopMonitorDataResponse", string(data)}, " ")
}
