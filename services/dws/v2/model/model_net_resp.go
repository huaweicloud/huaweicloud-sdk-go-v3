package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NetResp struct {

	// **参数解释**： 虚拟集群ID。 **取值范围**： 不涉及。
	VirtualClusterId *int32 `json:"virtual_cluster_id,omitempty"`

	// **参数解释**： 查询时间。 **取值范围**： 不涉及。
	Ctime *int64 `json:"ctime,omitempty"`

	// **参数解释**： 主机ID。 **取值范围**： 不涉及。
	HostId *int32 `json:"host_id,omitempty"`

	// **参数解释**： 主机名称。 **取值范围**： 不涉及。
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 网卡状态。 **取值范围**： true：代表up。 false：代表down。
	Up *bool `json:"up,omitempty"`

	// **参数解释**： 网卡速度(Mbps)。 **取值范围**： 不涉及。
	Speed *int64 `json:"speed,omitempty"`

	// **参数解释**： 接收包数(个)。 **取值范围**： 不涉及。
	RecvPackets *int64 `json:"recv_packets,omitempty"`

	// **参数解释**： 发送包数(个)。 **取值范围**： 不涉及。
	SendPackets *int64 `json:"send_packets,omitempty"`

	// **参数解释**： 接收丢包数(个)。 **取值范围**： 不涉及。
	RecvDrop *int64 `json:"recv_drop,omitempty"`

	// **参数解释**： 接收速率(KB/s)。 **取值范围**： 不涉及。
	RecvRate *float64 `json:"recv_rate,omitempty"`

	// **参数解释**： 发送速率(KB/s)。 **取值范围**： 不涉及。
	SendRate *float64 `json:"send_rate,omitempty"`

	// **参数解释**： 网络速率(KB/s)。 **取值范围**： 不涉及。
	IoRate *float64 `json:"io_rate,omitempty"`

	// **参数解释**： 网卡名称。 **取值范围**： 不涉及。
	InterfaceName *string `json:"interface_name,omitempty"`
}

func (o NetResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NetResp struct{}"
	}

	return strings.Join([]string{"NetResp", string(data)}, " ")
}
