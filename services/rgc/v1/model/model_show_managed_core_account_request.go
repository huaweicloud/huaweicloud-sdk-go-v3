package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowManagedCoreAccountRequest Request Object
type ShowManagedCoreAccountRequest struct {

	// 账号类型。包括LOGGING，SECURITY和PRIMARY账号。
	AccountType string `json:"account_type"`

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`
}

func (o ShowManagedCoreAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowManagedCoreAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowManagedCoreAccountRequest", string(data)}, " ")
}
