package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateManagedAccountRequest Request Object
type UpdateManagedAccountRequest struct {

	// 账号ID。
	ManagedAccountId string `json:"managed_account_id"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *CreateManagedAccountRequest `json:"body,omitempty"`
}

func (o UpdateManagedAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateManagedAccountRequest struct{}"
	}

	return strings.Join([]string{"UpdateManagedAccountRequest", string(data)}, " ")
}
