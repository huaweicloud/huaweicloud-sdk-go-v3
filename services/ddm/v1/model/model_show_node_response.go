package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowNodeResponse struct {

	// 节点状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 节点名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 节点id。
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 节点私有ip。
	PrivateIp *string `json:"private_ip,omitempty" xml:"private_ip"`

	// 节点浮动ip。
	FloatingIp *string `json:"floating_ip,omitempty" xml:"floating_ip"`

	// 虚机id。
	ServerId *string `json:"server_id,omitempty" xml:"server_id"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty" xml:"subnet_name"`

	// 数据盘id。
	DatavolumeId *string `json:"datavolume_id,omitempty" xml:"datavolume_id"`

	// 资源子网ip。
	ResSubnetIp *string `json:"res_subnet_ip,omitempty" xml:"res_subnet_ip"`

	// 系统盘id。
	SystemvolumeId *string `json:"systemvolume_id,omitempty" xml:"systemvolume_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNodeResponse struct{}"
	}

	return strings.Join([]string{"ShowNodeResponse", string(data)}, " ")
}
