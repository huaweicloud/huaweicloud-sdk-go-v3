package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnServerResponse Response Object
type CreateVpnServerResponse struct {
	VpnServer *CreateServerResponseBodyVpnServer `json:"vpn_server,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVpnServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnServerResponse struct{}"
	}

	return strings.Join([]string{"CreateVpnServerResponse", string(data)}, " ")
}
