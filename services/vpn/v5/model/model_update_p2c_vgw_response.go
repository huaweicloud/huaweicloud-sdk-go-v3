package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateP2cVgwResponse Response Object
type UpdateP2cVgwResponse struct {
	P2cVpnGateway *ResponseP2cVgw `json:"p2c_vpn_gateway,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateP2cVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateP2cVgwResponse struct{}"
	}

	return strings.Join([]string{"UpdateP2cVgwResponse", string(data)}, " ")
}
