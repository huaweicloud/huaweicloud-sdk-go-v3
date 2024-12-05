package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteConnectGatewayResponse Response Object
type DeleteConnectGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteConnectGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConnectGatewayResponse struct{}"
	}

	return strings.Join([]string{"DeleteConnectGatewayResponse", string(data)}, " ")
}
