package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateConnectGatewayRequestBody struct {
	ConnectGateway *CreateConnectGateway `json:"connect_gateway"`
}

func (o CreateConnectGatewayRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateConnectGatewayRequestBody struct{}"
	}

	return strings.Join([]string{"CreateConnectGatewayRequestBody", string(data)}, " ")
}
