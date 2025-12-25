package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FilterConnectionType **参数解释**: 连接类型 - FILTER 过滤 - INPUT 输入 - OUTPUT 输出  **约束限制** 不涉及 **取值范围**: - FILTER - INPUT - OUTPUT  **默认值** 不涉及
type FilterConnectionType struct {
	value string
}

type FilterConnectionTypeEnum struct {
	FILTER FilterConnectionType
	INPUT  FilterConnectionType
	OUTPUT FilterConnectionType
}

func GetFilterConnectionTypeEnum() FilterConnectionTypeEnum {
	return FilterConnectionTypeEnum{
		FILTER: FilterConnectionType{
			value: "FILTER",
		},
		INPUT: FilterConnectionType{
			value: "INPUT",
		},
		OUTPUT: FilterConnectionType{
			value: "OUTPUT",
		},
	}
}

func (c FilterConnectionType) Value() string {
	return c.value
}

func (c FilterConnectionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FilterConnectionType) UnmarshalJSON(b []byte) error {
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
