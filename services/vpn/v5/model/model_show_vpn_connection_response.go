package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnConnectionResponse Response Object
type ShowVpnConnectionResponse struct {
	VpnConnection *ResponseVpnConnection `json:"vpn_connection,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowVpnConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnConnectionResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnConnectionResponse", string(data)}, " ")
}
