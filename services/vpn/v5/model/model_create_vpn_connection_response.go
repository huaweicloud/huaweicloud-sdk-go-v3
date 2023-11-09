package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnConnectionResponse Response Object
type CreateVpnConnectionResponse struct {
	VpnConnection *CreateResponseVpnConnection `json:"vpn_connection,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVpnConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnConnectionResponse struct{}"
	}

	return strings.Join([]string{"CreateVpnConnectionResponse", string(data)}, " ")
}
