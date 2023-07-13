package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExternalGatewayInfoOption
type ExternalGatewayInfoOption struct {

	// 外部网络的ID。 外部网络的信息请通过GET /v2.0/networks?router:external=True或neutron net-external-list方式查询。
	NetworkId *string `json:"network_id,omitempty"`
}

func (o ExternalGatewayInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExternalGatewayInfoOption struct{}"
	}

	return strings.Join([]string{"ExternalGatewayInfoOption", string(data)}, " ")
}
