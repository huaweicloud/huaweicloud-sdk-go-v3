package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VirtualPortResponse IPv6 IP绑定带宽请求体。
type VirtualPortResponse struct {

	// IPv6带宽ID。
	Ipv6BandwidthId *string `json:"ipv6_bandwidth_id,omitempty"`

	// IPv6虚拟IP或私网IP ID。
	PortId *string `json:"port_id,omitempty"`
}

func (o VirtualPortResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VirtualPortResponse struct{}"
	}

	return strings.Join([]string{"VirtualPortResponse", string(data)}, " ")
}
