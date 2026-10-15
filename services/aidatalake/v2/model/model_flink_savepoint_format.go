package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkSavepointFormat **参数解释**：Savepoint格式类型。 **约束限制**：不涉及。 **取值范围**：   - NATIVE：原生格式。   - CANONICAL：通用格式。 **默认取值**：NATIVE。
type FlinkSavepointFormat struct {
	value string
}

type FlinkSavepointFormatEnum struct {
	NATIVE    FlinkSavepointFormat
	CANONICAL FlinkSavepointFormat
}

func GetFlinkSavepointFormatEnum() FlinkSavepointFormatEnum {
	return FlinkSavepointFormatEnum{
		NATIVE: FlinkSavepointFormat{
			value: "NATIVE",
		},
		CANONICAL: FlinkSavepointFormat{
			value: "CANONICAL",
		},
	}
}

func (c FlinkSavepointFormat) Value() string {
	return c.value
}

func (c FlinkSavepointFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkSavepointFormat) UnmarshalJSON(b []byte) error {
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
