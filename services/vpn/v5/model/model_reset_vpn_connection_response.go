package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetVpnConnectionResponse Response Object
type ResetVpnConnectionResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ResetVpnConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVpnConnectionResponse struct{}"
	}

	return strings.Join([]string{"ResetVpnConnectionResponse", string(data)}, " ")
}
