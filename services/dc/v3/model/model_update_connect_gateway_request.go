package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectGatewayRequest Request Object
type UpdateConnectGatewayRequest struct {

	// 互联网关ID
	ConnectGatewayId string `json:"connect_gateway_id"`

	Body *UpdateConnectGatewayRequestBody `json:"body,omitempty"`
}

func (o UpdateConnectGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectGatewayRequest struct{}"
	}

	return strings.Join([]string{"UpdateConnectGatewayRequest", string(data)}, " ")
}
