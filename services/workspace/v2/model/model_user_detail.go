package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserDetail struct {

	// 用户描述。
	Description *string `json:"description,omitempty"`

	// 用户id。
	Id *string `json:"id,omitempty"`

	// 桌面用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 用户sid。
	ObjectSid *string `json:"object_sid,omitempty"`

	// 登录名(windows以前版本)。
	SamAccountName *string `json:"sam_account_name,omitempty"`

	// 用户登录名。
	UserPrincipalName *string `json:"user_principal_name,omitempty"`

	// 全名。
	FullName *string `json:"full_name,omitempty"`

	// 用户在域树上的唯一位置。
	DistinguishedName *string `json:"distinguished_name,omitempty"`

	// 帐号类型(0：用户；1：用户组)。
	AccountType *int32 `json:"account_type,omitempty"`

	// UTC时间毫秒数对应的字符。
	WhenCreated *string `json:"when_created,omitempty"`

	// 账号有效期最后一天对应的UTC时间，以毫秒为单位。
	AccountExpires *int64 `json:"account_expires,omitempty"`

	// 账户是否过期，true表示过期，false表示未过期。
	UserExpired *bool `json:"user_expired,omitempty"`

	// 账户是否被锁定，true表示被锁定，false表示未锁定。
	Locked *bool `json:"locked,omitempty"`

	// 是否允许修改密码，true表示允许修改密码，false表示不允许。
	EnabledChangePassword *bool `json:"enabled_change_password,omitempty"`

	// 密码是否永不过期，true表示密码永不过期，false表示密码会过期。
	PasswordNeverExpired *bool `json:"password_never_expired,omitempty"`

	// 下次登录是否需要重置密码，true表示需要重置密码，false表示不需要。
	NextLoginChangePassword *bool `json:"next_login_change_password,omitempty"`

	// 账户是否禁用，true表示被禁用，false表示未禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// 加入的组列表。
	GroupNames *[]string `json:"group_names,omitempty"`

	// 用户绑定桌面云总数。
	TotalDesktops *int32 `json:"total_desktops,omitempty"`
}

func (o UserDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDetail struct{}"
	}

	return strings.Join([]string{"UserDetail", string(data)}, " ")
}
