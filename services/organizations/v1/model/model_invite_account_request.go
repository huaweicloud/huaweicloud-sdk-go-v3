package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InviteAccountRequest Request Object
type InviteAccountRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *InviteAccountReqBody `json:"body,omitempty"`
}

func (o InviteAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InviteAccountRequest struct{}"
	}

	return strings.Join([]string{"InviteAccountRequest", string(data)}, " ")
}
