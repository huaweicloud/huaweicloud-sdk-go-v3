package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateP2cVgwResponse Response Object
type CreateP2cVgwResponse struct {
	P2cVpnGateway *ResponseP2cVgw `json:"p2c_vpn_gateway,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateP2cVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateP2cVgwResponse struct{}"
	}

	return strings.Join([]string{"CreateP2cVgwResponse", string(data)}, " ")
}
