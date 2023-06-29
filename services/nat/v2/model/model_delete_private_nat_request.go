package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePrivateNatRequest Request Object
type DeletePrivateNatRequest struct {

	// 私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`
}

func (o DeletePrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePrivateNatRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivateNatRequest", string(data)}, " ")
}
