package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateUserReqBody 创建用户请求体
type CreateUserReqBody struct {

	// 用户的地址信息列表
	Addresses *[]AddressDto `json:"addresses,omitempty"`

	// 用户的显示名称
	DisplayName string `json:"display_name"`

	// 用户的电子邮箱信息列表
	Emails []EmailDto `json:"emails"`

	// 用户的地理区域或位置信息
	Locale *string `json:"locale,omitempty"`

	Name *NameDto `json:"name"`

	// 用户昵称
	Nickname *string `json:"nickname,omitempty"`

	// 用户的电话号码信息列表
	PhoneNumbers *[]PhoneNumberDto `json:"phone_numbers,omitempty"`

	// 用户语言首选项
	PreferredLanguage *string `json:"preferred_language,omitempty"`

	// 与用户关联的URL
	ProfileUrl *string `json:"profile_url,omitempty"`

	// 用户时区
	Timezone *string `json:"timezone,omitempty"`

	// 用户头衔
	Title *string `json:"title,omitempty"`

	// 用户名，用于标识用户的唯一字符串
	UserName string `json:"user_name"`

	// 用户类型
	UserType *string `json:"user_type,omitempty"`

	// 初始化密码方式，一次性密码/邮箱
	PasswordMode CreateUserReqBodyPasswordMode `json:"password_mode"`

	Enterprise *EnterpriseDto `json:"enterprise,omitempty"`
}

func (o CreateUserReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserReqBody struct{}"
	}

	return strings.Join([]string{"CreateUserReqBody", string(data)}, " ")
}

type CreateUserReqBodyPasswordMode struct {
	value string
}

type CreateUserReqBodyPasswordModeEnum struct {
	OTP   CreateUserReqBodyPasswordMode
	EMAIL CreateUserReqBodyPasswordMode
}

func GetCreateUserReqBodyPasswordModeEnum() CreateUserReqBodyPasswordModeEnum {
	return CreateUserReqBodyPasswordModeEnum{
		OTP: CreateUserReqBodyPasswordMode{
			value: "OTP",
		},
		EMAIL: CreateUserReqBodyPasswordMode{
			value: "EMAIL",
		},
	}
}

func (c CreateUserReqBodyPasswordMode) Value() string {
	return c.value
}

func (c CreateUserReqBodyPasswordMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserReqBodyPasswordMode) UnmarshalJSON(b []byte) error {
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
