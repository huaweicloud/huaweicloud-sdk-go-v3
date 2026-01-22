package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserGroupPerms struct {

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消费组权限。 **约束限制**： 不涉及。 **取值范围**： - SUB：拥有订阅权限。 - DENY：无权限。  **默认取值**： 不涉及。
	Perm *UserGroupPermsPerm `json:"perm,omitempty"`
}

func (o UserGroupPerms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserGroupPerms struct{}"
	}

	return strings.Join([]string{"UserGroupPerms", string(data)}, " ")
}

type UserGroupPermsPerm struct {
	value string
}

type UserGroupPermsPermEnum struct {
	SUB  UserGroupPermsPerm
	DENY UserGroupPermsPerm
}

func GetUserGroupPermsPermEnum() UserGroupPermsPermEnum {
	return UserGroupPermsPermEnum{
		SUB: UserGroupPermsPerm{
			value: "SUB",
		},
		DENY: UserGroupPermsPerm{
			value: "DENY",
		},
	}
}

func (c UserGroupPermsPerm) Value() string {
	return c.value
}

func (c UserGroupPermsPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserGroupPermsPerm) UnmarshalJSON(b []byte) error {
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
