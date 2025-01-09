package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNatGatewaysResponse Response Object
type ListNatGatewaysResponse struct {

	// 网关实例信息列表。
	NatGateways    *[]NatGateway `json:"nat_gateways,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListNatGatewaysResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNatGatewaysResponse struct{}"
	}

	return strings.Join([]string{"ListNatGatewaysResponse", string(data)}, " ")
}
