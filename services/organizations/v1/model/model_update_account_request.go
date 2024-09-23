package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAccountRequest Request Object
type UpdateAccountRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	Body *UpdateAccountReqBody `json:"body,omitempty"`
}

func (o UpdateAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAccountRequest struct{}"
	}

	return strings.Join([]string{"UpdateAccountRequest", string(data)}, " ")
}
