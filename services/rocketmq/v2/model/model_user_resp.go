package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserResp struct {

	// **参数解释**： 用户名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessKey *string `json:"access_key,omitempty"`

	// **参数解释**： 密钥。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SecretKey *string `json:"secret_key,omitempty"`

	// **参数解释**： IP白名单。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	WhiteRemoteAddress *string `json:"white_remote_address,omitempty"`

	// **参数解释**： 是否为管理员。 **约束限制**： 不涉及。 **取值范围**： - true：是管理员。 - false：不是管理员。 **默认取值**： 不涉及。
	Admin *bool `json:"admin,omitempty"`

	// **参数解释**： 默认的主题权限。 **约束限制**： 不涉及。 **取值范围**： - PUB：拥有发布权限。 - SUB：拥有订阅权限。 - PUB|SUB：拥有发布订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	DefaultTopicPerm *UserRespDefaultTopicPerm `json:"default_topic_perm,omitempty"`

	// **参数解释**： 默认的消费组权限。 **约束限制**： 不涉及。 **取值范围**： - SUB：拥有订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	DefaultGroupPerm *UserRespDefaultGroupPerm `json:"default_group_perm,omitempty"`

	// **参数解释**： 特殊的主题权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicPerms *[]UserTopicPerms `json:"topic_perms,omitempty"`

	// **参数解释**： 特殊的消费组权限。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	GroupPerms *[]UserRespGroupPerms `json:"group_perms,omitempty"`
}

func (o UserResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserResp struct{}"
	}

	return strings.Join([]string{"UserResp", string(data)}, " ")
}

type UserRespDefaultTopicPerm struct {
	value string
}

type UserRespDefaultTopicPermEnum struct {
	PUB     UserRespDefaultTopicPerm
	SUB     UserRespDefaultTopicPerm
	PUB_SUB UserRespDefaultTopicPerm
	DENY    UserRespDefaultTopicPerm
}

func GetUserRespDefaultTopicPermEnum() UserRespDefaultTopicPermEnum {
	return UserRespDefaultTopicPermEnum{
		PUB: UserRespDefaultTopicPerm{
			value: "PUB",
		},
		SUB: UserRespDefaultTopicPerm{
			value: "SUB",
		},
		PUB_SUB: UserRespDefaultTopicPerm{
			value: "PUB|SUB",
		},
		DENY: UserRespDefaultTopicPerm{
			value: "DENY",
		},
	}
}

func (c UserRespDefaultTopicPerm) Value() string {
	return c.value
}

func (c UserRespDefaultTopicPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRespDefaultTopicPerm) UnmarshalJSON(b []byte) error {
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

type UserRespDefaultGroupPerm struct {
	value string
}

type UserRespDefaultGroupPermEnum struct {
	SUB  UserRespDefaultGroupPerm
	DENY UserRespDefaultGroupPerm
}

func GetUserRespDefaultGroupPermEnum() UserRespDefaultGroupPermEnum {
	return UserRespDefaultGroupPermEnum{
		SUB: UserRespDefaultGroupPerm{
			value: "SUB",
		},
		DENY: UserRespDefaultGroupPerm{
			value: "DENY",
		},
	}
}

func (c UserRespDefaultGroupPerm) Value() string {
	return c.value
}

func (c UserRespDefaultGroupPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRespDefaultGroupPerm) UnmarshalJSON(b []byte) error {
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
