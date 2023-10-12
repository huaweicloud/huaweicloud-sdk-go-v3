package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountInfo 用户详细信息
type AccountInfo struct {

	// 用户(组)
	Account string `json:"account"`

	AccountType *AccountTypeEnum `json:"account_type"`

	// 域名城
	Domain *string `json:"domain,omitempty"`

	// 邮箱
	Email *string `json:"email,omitempty"`

	// 手机
	TelephoneNumber *string `json:"telephone_number,omitempty"`
}

func (o AccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountInfo struct{}"
	}

	return strings.Join([]string{"AccountInfo", string(data)}, " ")
}
