package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPrivateNatRequest Request Object
type ShowPrivateNatRequest struct {

	// 私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`
}

func (o ShowPrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPrivateNatRequest struct{}"
	}

	return strings.Join([]string{"ShowPrivateNatRequest", string(data)}, " ")
}
