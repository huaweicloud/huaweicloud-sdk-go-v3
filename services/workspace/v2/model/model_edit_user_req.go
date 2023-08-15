package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EditUserReq struct {

	// 用户描述。
	Description *string `json:"description,omitempty"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *EditUserReqActiveType `json:"active_type,omitempty"`

	// 账户过期时间，0表示永远不过期。
	AccountExpires *string `json:"account_expires,omitempty"`

	// 是否允许修改密码，true表示允许，false表示不允许。
	EnableChangePassword *bool `json:"enable_change_password,omitempty"`

	// 下次登录是否需要重置密码，true表示需要重置密码，false表示不需要。
	NextLoginChangePassword *bool `json:"next_login_change_password,omitempty"`

	// 密码是否永不过期，true表示密码永不过期，false表示密码会过期。
	PasswordNeverExpired *bool `json:"password_never_expired,omitempty"`

	// 账户是否禁用，true表示被禁用，false表示未禁用。
	Disabled *bool `json:"disabled,omitempty"`
}

func (o EditUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditUserReq struct{}"
	}

	return strings.Join([]string{"EditUserReq", string(data)}, " ")
}

type EditUserReqActiveType struct {
	value string
}

type EditUserReqActiveTypeEnum struct {
	USER_ACTIVATE  EditUserReqActiveType
	ADMIN_ACTIVATE EditUserReqActiveType
}

func GetEditUserReqActiveTypeEnum() EditUserReqActiveTypeEnum {
	return EditUserReqActiveTypeEnum{
		USER_ACTIVATE: EditUserReqActiveType{
			value: "USER_ACTIVATE",
		},
		ADMIN_ACTIVATE: EditUserReqActiveType{
			value: "ADMIN_ACTIVATE",
		},
	}
}

func (c EditUserReqActiveType) Value() string {
	return c.value
}

func (c EditUserReqActiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EditUserReqActiveType) UnmarshalJSON(b []byte) error {
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
