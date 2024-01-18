package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AccountInfo 用户详细信息。
type AccountInfo struct {

	// 用户ID(或用户组ID)，根据 account_type 参数决定值类型。 对于用户组类型，必须传入用户组ID。 `USER` - 用户ID。 `USER_GROUP` - 用户组ID。
	Id *string `json:"id,omitempty"`

	// 用户名(或用户组名)，根据 account_type 参数决定值类型。 `USER` - 用户名。 `USER_GROUP` - 用户组名。
	Account string `json:"account"`

	AccountType *AccountTypeEnum `json:"account_type"`

	// 域名城。
	Domain *string `json:"domain,omitempty"`

	// 邮箱。
	Email *string `json:"email,omitempty"`

	// 手机。
	TelephoneNumber *string `json:"telephone_number,omitempty"`

	PlatformType *PlatformTypeEnum `json:"platform_type,omitempty"`
}

func (o AccountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountInfo struct{}"
	}

	return strings.Join([]string{"AccountInfo", string(data)}, " ")
}
