package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConnectGatewayResponse Response Object
type UpdateConnectGatewayResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateConnectGatewayResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectGatewayResponse struct{}"
	}

	return strings.Join([]string{"UpdateConnectGatewayResponse", string(data)}, " ")
}
