package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type User struct {

	// 用户ID。
	Id *string `json:"id,omitempty"`

	// 用户ID。
	Sid *string `json:"sid,omitempty"`

	// 桌面用户名。
	UserName *string `json:"user_name,omitempty"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 用户绑定桌面云总数。
	TotalDesktops *int32 `json:"total_desktops,omitempty"`

	// 手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *UserActiveType `json:"active_type,omitempty"`

	// 是不是预创建的用户。
	IsPreUser *bool `json:"is_pre_user,omitempty"`

	// 账户过期时间，0表示永远不过期。
	AccountExpires *string `json:"account_expires,omitempty"`

	// 密码是否永不过期，true表示密码永不过期，false表示密码会过期。
	PasswordNeverExpired *bool `json:"password_never_expired,omitempty"`

	// 账号是否过期，true表示已过期，false表示未过期。
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

	// 用户是否订阅协同，true表示已订阅，false表示未订阅。
	ShareSpaceSubscription *bool `json:"share_space_subscription,omitempty"`

	// 用户已绑定协同桌面数。
	ShareSpaceDesktops *int32 `json:"share_space_desktops,omitempty"`

	// 加入的组列表。
	GroupNames *[]string `json:"group_names,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 用户信息映射，包含用户的服务等级、操作模式和类型。
	UserInfoMap *string `json:"user_info_map,omitempty"`

	// 域。
	Domain *string `json:"domain,omitempty"`

	// 当前用户是否存在用户证书。
	CertificateStatus *bool `json:"certificate_status,omitempty"`
}

func (o User) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "User struct{}"
	}

	return strings.Join([]string{"User", string(data)}, " ")
}

type UserActiveType struct {
	value string
}

type UserActiveTypeEnum struct {
	USER_ACTIVATE  UserActiveType
	ADMIN_ACTIVATE UserActiveType
}

func GetUserActiveTypeEnum() UserActiveTypeEnum {
	return UserActiveTypeEnum{
		USER_ACTIVATE: UserActiveType{
			value: "USER_ACTIVATE",
		},
		ADMIN_ACTIVATE: UserActiveType{
			value: "ADMIN_ACTIVATE",
		},
	}
}

func (c UserActiveType) Value() string {
	return c.value
}

func (c UserActiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserActiveType) UnmarshalJSON(b []byte) error {
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
