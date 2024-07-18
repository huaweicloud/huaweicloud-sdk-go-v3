package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVpnUserResponse Response Object
type CreateVpnUserResponse struct {
	User *CreateVpnUserResponseBodyUser `json:"user,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVpnUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVpnUserResponse struct{}"
	}

	return strings.Join([]string{"CreateVpnUserResponse", string(data)}, " ")
}
