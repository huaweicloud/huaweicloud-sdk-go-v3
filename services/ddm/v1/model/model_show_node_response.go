package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNodeResponse struct {
	// 节点状态。

	Status *string `json:"status,omitempty"`
	// 节点名称。

	Name *string `json:"name,omitempty"`
	// 节点id。

	NodeId *string `json:"node_id,omitempty"`
	// 节点私有ip。

	PrivateIp *string `json:"private_ip,omitempty"`
	// 节点浮动ip。

	FloatingIp *string `json:"floating_ip,omitempty"`
	// 虚机id。

	ServerId *string `json:"server_id,omitempty"`
	// 子网名称。

	SubnetName *string `json:"subnet_name,omitempty"`
	// 数据盘id。

	DatavolumeId *string `json:"datavolume_id,omitempty"`
	// 资源子网ip。

	ResSubnetIp *string `json:"res_subnet_ip,omitempty"`
	// 系统盘id。

	SystemvolumeId *string `json:"systemvolume_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeResponse", string(data)}, " ")
}
