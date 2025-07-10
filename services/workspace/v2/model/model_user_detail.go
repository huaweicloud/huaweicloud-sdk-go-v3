package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

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

	// 手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *UserDetailActiveType `json:"active_type,omitempty"`

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

	// 账号类型(0：用户；1：用户组)。
	AccountType *int32 `json:"account_type,omitempty"`

	// UTC时间毫秒数对应的字符，格式为：yyyy-MM-ddTHH:mm:ss.SSSZ。
	WhenCreated *string `json:"when_created,omitempty"`

	// 账号有效期最后一天对应的UTC时间，以毫秒为单位。
	AccountExpires *int64 `json:"account_expires,omitempty"`

	// 是否是预创建的用户，true表示是预创建用户，false表示不是预创建用户。
	IsPreUser *bool `json:"is_pre_user,omitempty"`

	// 账户是否过期，true表示过期，false表示未过期。
	UserExpired *bool `json:"user_expired,omitempty"`

	// 账户是否被锁定，true表示被锁定，false表示未锁定。
	Locked *bool `json:"locked,omitempty"`

	// 是否允许修改密码，true表示允许修改密码，false表示不允许。
	EnableChangePassword *bool `json:"enable_change_password,omitempty"`

	// 密码是否永不过期，true表示密码永不过期，false表示密码会过期。
	PasswordNeverExpired *bool `json:"password_never_expired,omitempty"`

	// 下次登录是否需要重置密码，true表示需要重置密码，false表示不需要。
	NextLoginChangePassword *bool `json:"next_login_change_password,omitempty"`

	// 账户是否禁用，true表示被禁用，false表示未禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// 加入的组列表。
	GroupNames *[]string `json:"group_names,omitempty"`

	// 用户是否订阅协同，true表示已订阅，false表示未订阅。
	ShareSpaceSubscription *bool `json:"share_space_subscription,omitempty"`

	// 用户已绑定协同桌面数。
	ShareSpaceDesktops *int32 `json:"share_space_desktops,omitempty"`

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

type UserDetailActiveType struct {
	value string
}

type UserDetailActiveTypeEnum struct {
	ADMIN_ACTIVATE UserDetailActiveType
	USER_ACTIVATE  UserDetailActiveType
}

func GetUserDetailActiveTypeEnum() UserDetailActiveTypeEnum {
	return UserDetailActiveTypeEnum{
		ADMIN_ACTIVATE: UserDetailActiveType{
			value: "ADMIN_ACTIVATE",
		},
		USER_ACTIVATE: UserDetailActiveType{
			value: "USER_ACTIVATE",
		},
	}
}

func (c UserDetailActiveType) Value() string {
	return c.value
}

func (c UserDetailActiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserDetailActiveType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
