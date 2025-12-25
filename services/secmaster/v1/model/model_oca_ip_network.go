package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OcaIpNetwork 线下机房网络信息
type OcaIpNetwork struct {

	// 外网或内网 true：外网 false: 内网
	IsPublic bool `json:"is_public"`

	// 网络分区：OM/PSZ/DMZ
	Partition *string `json:"partition,omitempty"`

	// 网络平面（线下有自己的代号）
	Plane *string `json:"plane,omitempty"`

	// 虚拟网络ID
	VxlanId *string `json:"vxlan_id,omitempty"`
}

func (o OcaIpNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OcaIpNetwork struct{}"
	}

	return strings.Join([]string{"OcaIpNetwork", string(data)}, " ")
}
