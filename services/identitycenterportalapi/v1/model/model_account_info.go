package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountInfo Provides information about your account.
type AccountInfo struct {

	// 分配给用户的账号的全局唯一标识符（ID）
	AccountId *string `json:"account_id,omitempty"`

	// 分配给用户的账号的名称
	AccountName *string `json:"account_name,omitempty"`

	// 分配给用户的账号的电子邮箱地址
	EmailAddress *string `json:"email_address,omitempty"`
}

func (o AccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountInfo struct{}"
	}

	return strings.Join([]string{"AccountInfo", string(data)}, " ")
}
