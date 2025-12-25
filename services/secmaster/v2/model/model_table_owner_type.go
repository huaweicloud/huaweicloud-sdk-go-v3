package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableOwnerType **参数解释**: 拥有者类型 - SYSTEM 系统 - USER 用户 - SYSTEM_ALLOWED_DELETE 系统可删除 - USER_ALLOWED_DELETE 用户可删除  **约束限制** 不涉及 **取值范围**: - SYSTEM - USER - SYSTEM_ALLOWED_DELETE - USER_ALLOWED_DELETE  **默认值** 不涉及
type TableOwnerType struct {
	value string
}

type TableOwnerTypeEnum struct {
	SYSTEM                TableOwnerType
	USER                  TableOwnerType
	SYSTEM_ALLOWED_DELETE TableOwnerType
	USER_ALLOWED_DELETE   TableOwnerType
}

func GetTableOwnerTypeEnum() TableOwnerTypeEnum {
	return TableOwnerTypeEnum{
		SYSTEM: TableOwnerType{
			value: "SYSTEM",
		},
		USER: TableOwnerType{
			value: "USER",
		},
		SYSTEM_ALLOWED_DELETE: TableOwnerType{
			value: "SYSTEM_ALLOWED_DELETE",
		},
		USER_ALLOWED_DELETE: TableOwnerType{
			value: "USER_ALLOWED_DELETE",
		},
	}
}

func (c TableOwnerType) Value() string {
	return c.value
}

func (c TableOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableOwnerType) UnmarshalJSON(b []byte) error {
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
