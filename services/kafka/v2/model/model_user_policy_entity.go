package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserPolicyEntity struct {

	// **参数解释**： 资源类型。 **约束限制**： 不涉及。 **取值范围**： - TOPIC：表示资源类型为Topic。 **默认取值**： TOPIC
	ResourceType *string `json:"resource_type,omitempty"`

	// **参数解释**： 资源名称。 **约束限制**： 不涉及。 **取值范围**： - 已有的Topic名称。 - 符合Topic名称规则的前缀。 - “*”表示匹配所有的Topic。 **默认取值**： 不涉及。
	ResourceName *string `json:"resource_name,omitempty"`

	// **参数解释**： 匹配方式。 **约束限制**： 不涉及。 **取值范围**： - LITERAL：完全匹配。 - PREFIXED：前缀匹配。 **默认取值**： 不涉及。
	PatternType *UserPolicyEntityPatternType `json:"pattern_type,omitempty"`

	// **参数解释**： 权限类型。 **约束限制**： 不涉及。 **取值范围**： - all：拥有发布、订阅权限。 - pub：拥有发布权限。 - sub：拥有订阅权限。 **默认取值**： 不涉及。
	AccessPolicy *UserPolicyEntityAccessPolicy `json:"access_policy,omitempty"`
}

func (o UserPolicyEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserPolicyEntity struct{}"
	}

	return strings.Join([]string{"UserPolicyEntity", string(data)}, " ")
}

type UserPolicyEntityPatternType struct {
	value string
}

type UserPolicyEntityPatternTypeEnum struct {
	LITERAL  UserPolicyEntityPatternType
	PREFIXED UserPolicyEntityPatternType
}

func GetUserPolicyEntityPatternTypeEnum() UserPolicyEntityPatternTypeEnum {
	return UserPolicyEntityPatternTypeEnum{
		LITERAL: UserPolicyEntityPatternType{
			value: "LITERAL",
		},
		PREFIXED: UserPolicyEntityPatternType{
			value: "PREFIXED",
		},
	}
}

func (c UserPolicyEntityPatternType) Value() string {
	return c.value
}

func (c UserPolicyEntityPatternType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserPolicyEntityPatternType) UnmarshalJSON(b []byte) error {
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

type UserPolicyEntityAccessPolicy struct {
	value string
}

type UserPolicyEntityAccessPolicyEnum struct {
	ALL UserPolicyEntityAccessPolicy
	PUB UserPolicyEntityAccessPolicy
	SUB UserPolicyEntityAccessPolicy
}

func GetUserPolicyEntityAccessPolicyEnum() UserPolicyEntityAccessPolicyEnum {
	return UserPolicyEntityAccessPolicyEnum{
		ALL: UserPolicyEntityAccessPolicy{
			value: "all",
		},
		PUB: UserPolicyEntityAccessPolicy{
			value: "pub",
		},
		SUB: UserPolicyEntityAccessPolicy{
			value: "sub",
		},
	}
}

func (c UserPolicyEntityAccessPolicy) Value() string {
	return c.value
}

func (c UserPolicyEntityAccessPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserPolicyEntityAccessPolicy) UnmarshalJSON(b []byte) error {
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
