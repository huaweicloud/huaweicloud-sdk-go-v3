package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UserRespGroupPerms struct {

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消费组权限。 **约束限制**： 不涉及。 **取值范围**： - SUB：拥有订阅权限。 - DENY：无权限。 **默认取值**： 不涉及。
	Perm *UserRespGroupPermsPerm `json:"perm,omitempty"`
}

func (o UserRespGroupPerms) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserRespGroupPerms struct{}"
	}

	return strings.Join([]string{"UserRespGroupPerms", string(data)}, " ")
}

type UserRespGroupPermsPerm struct {
	value string
}

type UserRespGroupPermsPermEnum struct {
	SUB  UserRespGroupPermsPerm
	DENY UserRespGroupPermsPerm
}

func GetUserRespGroupPermsPermEnum() UserRespGroupPermsPermEnum {
	return UserRespGroupPermsPermEnum{
		SUB: UserRespGroupPermsPerm{
			value: "SUB",
		},
		DENY: UserRespGroupPermsPerm{
			value: "DENY",
		},
	}
}

func (c UserRespGroupPermsPerm) Value() string {
	return c.value
}

func (c UserRespGroupPermsPerm) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UserRespGroupPermsPerm) UnmarshalJSON(b []byte) error {
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
