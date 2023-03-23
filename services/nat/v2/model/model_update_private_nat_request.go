package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdatePrivateNatRequest struct {

	// 私网NAT网关实例的ID。
	GatewayId string `json:"gateway_id"`

	Body *UpdatePrivateNatRequestBody `json:"body,omitempty"`
}

func (o UpdatePrivateNatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePrivateNatRequest struct{}"
	}

	return strings.Join([]string{"UpdatePrivateNatRequest", string(data)}, " ")
}
