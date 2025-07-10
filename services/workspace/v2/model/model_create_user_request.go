package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateUserRequest struct {

	// 用户名称。
	UserName string `json:"user_name"`

	// 用户邮箱。
	UserEmail *string `json:"user_email,omitempty"`

	// 账户过期时间，0表示永远不过期。时间格式：yyyy-MM-ddTHH:mm:ssZ或yyyy-MM-ddTHH:mm:ss.SSSZ。
	AccountExpires *string `json:"account_expires,omitempty"`

	// 激活类型，默认为用户激活。 * USER_ACTIVATE： 用户激活 * ADMIN_ACTIVATE： 管理员激活
	ActiveType *CreateUserRequestActiveType `json:"active_type,omitempty"`

	// 用户手机号。
	UserPhone *string `json:"user_phone,omitempty"`

	// 用户初始密码。管理员激活模式需要输入。
	Password *string `json:"password,omitempty"`

	// 是否允许用户更改密码，缺省值为true，后续此字段无效，创建时都为true。
	EnableChangePassword *bool `json:"enable_change_password,omitempty"`

	// 下次登录是否必须更改密码，缺省值为true。后续此字段无效，创建时都为true。
	NextLoginChangePassword *bool `json:"next_login_change_password,omitempty"`

	// 用户组的专有ID列表。
	GroupIds *[]string `json:"group_ids,omitempty"`

	// 用户描述，字符串长度区间[0, 255]。
	Description *string `json:"description,omitempty"`

	// 别名。
	AliasName *string `json:"alias_name,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 用户信息映射，包含用户的服务等级、操作模式和类型。
	UserInfoMap *string `json:"user_info_map,omitempty"`
}

func (o CreateUserRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserRequest struct{}"
	}

	return strings.Join([]string{"CreateUserRequest", string(data)}, " ")
}

type CreateUserRequestActiveType struct {
	value string
}

type CreateUserRequestActiveTypeEnum struct {
	USER_ACTIVATE  CreateUserRequestActiveType
	ADMIN_ACTIVATE CreateUserRequestActiveType
}

func GetCreateUserRequestActiveTypeEnum() CreateUserRequestActiveTypeEnum {
	return CreateUserRequestActiveTypeEnum{
		USER_ACTIVATE: CreateUserRequestActiveType{
			value: "USER_ACTIVATE",
		},
		ADMIN_ACTIVATE: CreateUserRequestActiveType{
			value: "ADMIN_ACTIVATE",
		},
	}
}

func (c CreateUserRequestActiveType) Value() string {
	return c.value
}

func (c CreateUserRequestActiveType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserRequestActiveType) UnmarshalJSON(b []byte) error {
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
