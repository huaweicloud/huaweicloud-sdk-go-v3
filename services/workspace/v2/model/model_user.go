package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type User struct {

	// 用户ID。
	Id *string `json:"id,omitempty"`

	// 桌面用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 用户绑定桌面云总数。
	TotalDesktops *int32 `json:"total_desktops,omitempty"`

	// 账户过期时间，0表示永远不过期。
	AccountExpires *string `json:"account_expires,omitempty"`

	// 密码是否永不过期，true表示密码永不过期，false表示密码会过期。
	PasswordNeverExpired *bool `json:"password_never_expired,omitempty"`

	// 帐号是否过期，true表示已过期，false表示未过期。
	AccountExpired *bool `json:"account_expired,omitempty"`

	// 是否允许修改密码，true表示允许，false表示不允许。
	EnableChangePassword *bool `json:"enable_change_password,omitempty"`

	// 下次登录是否需要重置密码，true表示需要重置密码，false表示不需要。
	NextLoginChangePassword *bool `json:"next_login_change_password,omitempty"`

	// 用户描述。
	Description *string `json:"description,omitempty"`

	// 账户是否被锁定，true表示被锁定，false表示未锁定。
	Locked *bool `json:"locked,omitempty"`

	// 账户是否禁用，true表示被禁用，false表示未禁用。
	Disabled *bool `json:"disabled,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}
