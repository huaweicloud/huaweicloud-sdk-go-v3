package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteP2cVpnGatewayJobResponse Response Object
type DeleteP2cVpnGatewayJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteP2cVpnGatewayJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteP2cVpnGatewayJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteP2cVpnGatewayJobResponse", string(data)}, " ")
}
