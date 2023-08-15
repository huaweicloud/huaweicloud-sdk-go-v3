package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateUserResponse Response Object
type CreateUserResponse struct {

	// 用户名，只能英文字母开头，且由英文字母、数字、中划线、下划线组成，长度为7~64个字符。
	AccessKey *string `json:"access_key,omitempty"`

	// 密钥。 8-32个字符。 至少包含以下字符中的3种：   - 大写字母   - 小写字母   - 数字   - 特殊字符`~!@#$%^&*()-_=+\\\\|[{}];:\\'\\\",<.>/?密钥。 不能与名称或倒序的名称相同。
	SecretKey *string `json:"secret_key,omitempty"`

	// IP白名单。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// 是否为管理员。
	Admin *bool `json:"admin,omitempty"`

	// 默认的主题权限。
	DefaultTopicPerm *CreateUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	// 默认的消费组权限。
	DefaultGroupPerm *CreateUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty"`

	// 特殊的主题权限。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	// 特殊的消费组权限。
	GroupPerms     *[]UserGroupPerms `json:"group_perms,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserResponse struct{}"
	}

	return strings.Join([]string{"CreateUserResponse", string(data)}, " ")
}

type CreateUserResponseDefaultTopicPerm struct {
	value string
}

type CreateUserResponseDefaultTopicPermEnum struct {
	PUB     CreateUserResponseDefaultTopicPerm
	SUB     CreateUserResponseDefaultTopicPerm
	PUB_SUB CreateUserResponseDefaultTopicPerm
	DENY    CreateUserResponseDefaultTopicPerm
}

func GetCreateUserResponseDefaultTopicPermEnum() CreateUserResponseDefaultTopicPermEnum {
	return CreateUserResponseDefaultTopicPermEnum{
		PUB: CreateUserResponseDefaultTopicPerm{
			value: "PUB",
		},
		SUB: CreateUserResponseDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: CreateUserResponseDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: CreateUserResponseDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c CreateUserResponseDefaultTopicPerm) Value() string {
	return c.value
}

func (c CreateUserResponseDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserResponseDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type CreateUserResponseDefaultGroupPerm struct {
	value string
}

type CreateUserResponseDefaultGroupPermEnum struct {
	SUB  CreateUserResponseDefaultGroupPerm
	DENY CreateUserResponseDefaultGroupPerm
}

func GetCreateUserResponseDefaultGroupPermEnum() CreateUserResponseDefaultGroupPermEnum {
	return CreateUserResponseDefaultGroupPermEnum{
		SUB: CreateUserResponseDefaultGroupPerm{
			value: "SUB",
		},
		DENY: CreateUserResponseDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c CreateUserResponseDefaultGroupPerm) Value() string {
	return c.value
}

func (c CreateUserResponseDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateUserResponseDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
