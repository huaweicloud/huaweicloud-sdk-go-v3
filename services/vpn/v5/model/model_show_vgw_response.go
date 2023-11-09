package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVgwResponse Response Object
type ShowVgwResponse struct {
	VpnGateway *ResponseVpnGateway `json:"vpn_gateway,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVgwResponse struct{}"
	}

	return strings.Join([]string{"ShowVgwResponse", string(data)}, " ")
}
