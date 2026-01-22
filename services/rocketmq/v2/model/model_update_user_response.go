package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateUserResponse Response Object
type UpdateUserResponse struct {

	// **参数解释**： 用户名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessKey *string `json:"access_key,omitempty"`

	// **参数解释**： 密钥。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释**： IP白名单。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// **参数解释**： 是否为管理员。 **约束限制**： 不涉及。 **取值范围**： - true：是管理员。 - false：不是管理员。 **默认取值**： 不涉及。
	Admin *bool `json:"admin,omitempty"`

	// **参数解释**： 默认的主题权限。 **约束限制**： 不涉及。 **取值范围**： - PUB：拥有发布权限。 - SUB：拥有订阅权限。 - PUB|SUB：拥有发布订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	DefaultTopicPerm *UpdateUserResponseDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	// **参数解释**： 默认的消费组权限。 **约束限制**： 不涉及。 **取值范围**： - SUB：拥有订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	DefaultGroupPerm *UpdateUserResponseDefaultGroupPerm `json:"default_group_perm,omitempty"`

	// **参数解释**： 特殊的主题权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	// **参数解释**： 特殊的消费组权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupPerms     *[]UserRespGroupPerms `json:"group_perms,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o UpdateUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserResponse struct{}"
	}

	return strings.Join([]string{"UpdateUserResponse", string(data)}, " ")
}

type UpdateUserResponseDefaultTopicPerm struct {
	value string
}

type UpdateUserResponseDefaultTopicPermEnum struct {
	PUB     UpdateUserResponseDefaultTopicPerm
	SUB     UpdateUserResponseDefaultTopicPerm
	PUB_SUB UpdateUserResponseDefaultTopicPerm
	DENY    UpdateUserResponseDefaultTopicPerm
}

func GetUpdateUserResponseDefaultTopicPermEnum() UpdateUserResponseDefaultTopicPermEnum {
	return UpdateUserResponseDefaultTopicPermEnum{
		PUB: UpdateUserResponseDefaultTopicPerm{
			value: "PUB",
		},
		SUB: UpdateUserResponseDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: UpdateUserResponseDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: UpdateUserResponseDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c UpdateUserResponseDefaultTopicPerm) Value() string {
	return c.value
}

func (c UpdateUserResponseDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserResponseDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type UpdateUserResponseDefaultGroupPerm struct {
	value string
}

type UpdateUserResponseDefaultGroupPermEnum struct {
	SUB  UpdateUserResponseDefaultGroupPerm
	DENY UpdateUserResponseDefaultGroupPerm
}

func GetUpdateUserResponseDefaultGroupPermEnum() UpdateUserResponseDefaultGroupPermEnum {
	return UpdateUserResponseDefaultGroupPermEnum{
		SUB: UpdateUserResponseDefaultGroupPerm{
			value: "SUB",
		},
		DENY: UpdateUserResponseDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c UpdateUserResponseDefaultGroupPerm) Value() string {
	return c.value
}

func (c UpdateUserResponseDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateUserResponseDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
