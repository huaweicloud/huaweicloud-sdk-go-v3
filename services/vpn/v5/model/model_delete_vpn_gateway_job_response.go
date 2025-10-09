package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteVpnGatewayJobResponse Response Object
type DeleteVpnGatewayJobResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteVpnGatewayJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteVpnGatewayJobResponse struct{}"
	}

	return strings.Join([]string{"DeleteVpnGatewayJobResponse", string(data)}, " ")
}
