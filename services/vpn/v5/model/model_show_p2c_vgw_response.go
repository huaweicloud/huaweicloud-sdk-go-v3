package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowP2cVgwResponse Response Object
type ShowP2cVgwResponse struct {
	P2cVpnGateway *ShowResponseP2cVgw `json:"p2c_vpn_gateway,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowP2cVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowP2cVgwResponse struct{}"
	}

	return strings.Join([]string{"ShowP2cVgwResponse", string(data)}, " ")
}
