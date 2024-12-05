package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConnectGatewayRequestBody struct {
	ConnectGateway *UpdateConnectGateway `json:"connect_gateway,omitempty"`
}

func (o UpdateConnectGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateConnectGatewayRequestBody", string(data)}, " ")
}
