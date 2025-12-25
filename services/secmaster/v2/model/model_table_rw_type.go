package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TableRwType **参数解释**: 表读写类型 - READ_ONLY 只读 - READ_WRITE 读写 - WRITE_ONLY 只写  **约束限制** 不涉及 **取值范围**: - READ_ONLY - READ_WRITE - WRITE_ONLY  **默认值** 不涉及
type TableRwType struct {
	value string
}

type TableRwTypeEnum struct {
	READ_ONLY  TableRwType
	READ_WRITE TableRwType
	WRITE_ONLY TableRwType
}

func GetTableRwTypeEnum() TableRwTypeEnum {
	return TableRwTypeEnum{
		READ_ONLY: TableRwType{
			value: "READ_ONLY",
		},
		READ_WRITE: TableRwType{
			value: "READ_WRITE",
		},
		WRITE_ONLY: TableRwType{
			value: "WRITE_ONLY",
		},
	}
}

func (c TableRwType) Value() string {
	return c.value
}

func (c TableRwType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TableRwType) UnmarshalJSON(b []byte) error {
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
