package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConnectGatewayRequest Request Object
type ShowConnectGatewayRequest struct {

	// 互联网关ID
	ConnectGatewayId string `json:"connect_gateway_id"`
}

func (o ShowConnectGatewayRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConnectGatewayRequest struct{}"
	}

	return strings.Join([]string{"ShowConnectGatewayRequest", string(data)}, " ")
}
