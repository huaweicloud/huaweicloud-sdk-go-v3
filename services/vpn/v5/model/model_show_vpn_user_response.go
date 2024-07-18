package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnUserResponse Response Object
type ShowVpnUserResponse struct {
	User *VpnUser `json:"user,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowVpnUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnUserResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnUserResponse", string(data)}, " ")
}
