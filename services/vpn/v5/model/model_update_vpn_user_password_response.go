package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVpnUserPasswordResponse Response Object
type UpdateVpnUserPasswordResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVpnUserPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnUserPasswordResponse struct{}"
	}

	return strings.Join([]string{"UpdateVpnUserPasswordResponse", string(data)}, " ")
}
