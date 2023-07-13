package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNatGatewayRequestBody 更新公网NAT网关实例的请求体
type UpdateNatGatewayRequestBody struct {
	NatGateway *UpdateNatGatewayOption `json:"nat_gateway"`
}

func (o UpdateNatGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNatGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateNatGatewayRequestBody", string(data)}, " ")
}
