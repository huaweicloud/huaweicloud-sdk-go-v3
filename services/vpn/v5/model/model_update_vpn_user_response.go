package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnUserResponse Response Object
type UpdateVpnUserResponse struct {
	User *VpnUser `json:"user,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserResponse", string(data)}, " ")
}
