package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableEditType **参数解释**: 表编辑类型 - MODIFIABLE 可任意修改态 - APPENDED 追加态 (原来的内容锁定，只可追加) - LOCKED 锁定态 (完全锁定，不可修改)  **约束限制** 不涉及 **取值范围**: - MODIFIABLE - APPENDED - LOCKED  **默认值** 不涉及
type TableEditType struct {
	value string
}

type TableEditTypeEnum struct {
	MODIFIABLE TableEditType
	APPENDED   TableEditType
	LOCKED     TableEditType
}

func GetTableEditTypeEnum() TableEditTypeEnum {
	return TableEditTypeEnum{
		MODIFIABLE: TableEditType{
			value: "MODIFIABLE",
		},
		APPENDED: TableEditType{
			value: "APPENDED",
		},
		LOCKED: TableEditType{
			value: "LOCKED",
		},
	}
}

func (c TableEditType) Value() string {
	return c.value
}

func (c TableEditType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableEditType) UnmarshalJSON(b []byte) error {
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
