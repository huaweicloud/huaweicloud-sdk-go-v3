package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNatGatewayRequestBody 创建公网NAT网关实例的请求体。
type CreateNatGatewayRequestBody struct {
	NatGateway *CreateNatGatewayOption `json:"nat_gateway"`
}

func (o CreateNatGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNatGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateNatGatewayRequestBody", string(data)}, " ")
}
