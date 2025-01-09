package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountInfo 账户详细信息。
type AccountInfo struct {

	// 账户，账户的格式必须为:<i>账户(组)</i>的形式。
	Account string `json:"account"`

	// 域名(用户组必填，不填时使用默认值 local.com)。
	Domain *string `json:"domain,omitempty"`

	AccountType *AccountTypeEnum `json:"account_type"`

	PlatformType *PlatformTypeEnum `json:"platform_type,omitempty"`
}

func (o AccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountInfo struct{}"
	}

	return strings.Join([]string{"AccountInfo", string(data)}, " ")
}
