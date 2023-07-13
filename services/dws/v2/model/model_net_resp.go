package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetResp struct {

	// 虚拟集群ID。
	VirtualClusterId *int32 `json:"virtual_cluster_id,omitempty"`

	// 查询时间。
	Ctime *int64 `json:"ctime,omitempty"`

	// 主机ID。
	HostId *int32 `json:"host_id,omitempty"`

	// 主机名称。
	HostName *string `json:"host_name,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 网卡状态（true代表up/false代表down）。
	Up *bool `json:"up,omitempty"`

	// 网卡速度(Mbps)。
	Speed *int64 `json:"speed,omitempty"`

	// 接收包数(个)。
	RecvPackets *int64 `json:"recv_packets,omitempty"`

	// 发送包数(个)。
	SendPackets *int64 `json:"send_packets,omitempty"`

	// 接收丢包数(个)。
	RecvDrop *int64 `json:"recv_drop,omitempty"`

	// 接收速率(KB/s)。
	RecvRate *float64 `json:"recv_rate,omitempty"`

	// 发送速率(KB/s)。
	SendRate *float64 `json:"send_rate,omitempty"`

	// 网络速率(KB/s)。
	IoRate *float64 `json:"io_rate,omitempty"`

	// 网卡名称。
	InterfaceName *string `json:"interface_name,omitempty"`
}

func (o NetResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetResp struct{}"
	}

	return strings.Join([]string{"NetResp", string(data)}, " ")
}
