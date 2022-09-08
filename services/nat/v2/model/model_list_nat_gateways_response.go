package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListNatGatewaysResponse struct {

	// 查询公网NAT网关实例列表的响应体。 详见NatGateway字段说明。
	NatGateways    *[]NatGatewayResponseBody `json:"nat_gateways,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListNatGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewaysResponse", string(data)}, " ")
}
