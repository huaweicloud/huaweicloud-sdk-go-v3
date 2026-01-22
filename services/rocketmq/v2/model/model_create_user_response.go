package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateUserResponse Response Object
type CreateUserResponse struct {

	// **参数解释**： 用户名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessKey *string `json:"access_key,omitempty"`

	// **参数解释**： 密钥。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释**： IP白名单。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// **参数解释**： 是否为管理员。 **约束限制**： 不涉及。 **取值范围**： - true：是管理员。 - false：不是管理员。 **默认取值**： 不涉及。
	Admin *bool `json:"admin,omitempty"`

	// **参数解释**： 默认的主题权限。 **约束限制**： 不涉及。 **取值范围**： - PUB：拥有发布权限。 - SUB：拥有订阅权限。 - PUB|SUB：拥有发布订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	DefaultTopicPerm *CreateUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	// **参数解释**： 默认的消费组权限。 **约束限制**： 不涉及。 **取值范围**： - SUB：拥有订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	DefaultGroupPerm *CreateUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty"`

	// **参数解释**： 特殊的主题权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	// **参数解释**： 特殊的消费组权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupPerms     *[]UserRespGroupPerms `json:"group_perms,omitempty"`
	HttpStatusCode int                   `json:"-"`
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
