package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectGatewayRequest Request Object
type DeleteConnectGatewayRequest struct {

	// 互联网关ID
	ConnectGatewayId string `json:"connect_gateway_id"`
}

func (o DeleteConnectGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectGatewayRequest struct{}"
	}

	return strings.Join([]string{"DeleteConnectGatewayRequest", string(data)}, " ")
}
