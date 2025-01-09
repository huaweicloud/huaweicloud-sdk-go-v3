package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerMetricDataResponse Response Object
type ShowServerMetricDataResponse struct {

	// cpu使用率信息。
	CpuInfo *[]DataPoints `json:"cpu_info,omitempty"`

	// 内存使用率信息。
	MemoryInfo *[]DataPoints `json:"memory_info,omitempty"`

	// 磁盘使用率。
	DiskUtilInband *[]DataPoints `json:"disk_util_inband,omitempty"`

	// 磁盘读带宽。
	DiskReadBytesRate *[]DataPoints `json:"disk_read_bytes_rate,omitempty"`

	// 磁盘写带宽。
	DiskWriteBytesRate *[]DataPoints `json:"disk_write_bytes_rate,omitempty"`

	// 磁盘读IOPS。
	DiskReadRequestsRate *[]DataPoints `json:"disk_read_requests_rate,omitempty"`

	// 磁盘写IOPS。
	DiskWriteRequestsRate *[]DataPoints `json:"disk_write_requests_rate,omitempty"`

	// 带内网络流入速率。
	NetworkIncomingBytesRateInband *[]DataPoints `json:"network_incoming_bytes_rate_inband,omitempty"`

	// 带内网络流出速率。
	NetworkOutgoingBytesRateInband *[]DataPoints `json:"network_outgoing_bytes_rate_inband,omitempty"`

	// 带外网络流入速率。
	NetworkIncomingBytesAggregateRate *[]DataPoints `json:"network_incoming_bytes_aggregate_rate,omitempty"`

	// 带外网络流出速率。
	NetworkOutgoingBytesAggregateRate *[]DataPoints `json:"network_outgoing_bytes_aggregate_rate,omitempty"`

	// 网络连接数。
	NetworkVmConnections *[]DataPoints `json:"network_vm_connections,omitempty"`
	HttpStatusCode       int           `json:"-"`
}

func (o ShowServerMetricDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerMetricDataResponse struct{}"
	}

	return strings.Join([]string{"ShowServerMetricDataResponse", string(data)}, " ")
}
