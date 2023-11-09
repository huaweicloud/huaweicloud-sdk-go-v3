package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnConnectionResponse Response Object
type UpdateVpnConnectionResponse struct {
	VpnConnection *UpdateResponseVpnConnection `json:"vpn_connection,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnConnectionResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnConnectionResponse", string(data)}, " ")
}
