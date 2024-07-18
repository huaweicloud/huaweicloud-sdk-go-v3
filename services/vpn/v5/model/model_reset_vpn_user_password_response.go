package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetVpnUserPasswordResponse Response Object
type ResetVpnUserPasswordResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ResetVpnUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetVpnUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetVpnUserPasswordResponse", string(data)}, " ")
}
