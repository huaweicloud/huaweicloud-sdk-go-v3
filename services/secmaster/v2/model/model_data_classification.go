package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataClassification **参数解释**: 数据分类 - FACTUAL_DATA 事实数据 - DIMENSION_DATA 维度数据  **约束限制** 不涉及 **取值范围**: - FACTUAL_DATA - DIMENSION_DATA  **默认值** 不涉及
type DataClassification struct {
	value string
}

type DataClassificationEnum struct {
	FACTUAL_DATA   DataClassification
	DIMENSION_DATA DataClassification
}

func GetDataClassificationEnum() DataClassificationEnum {
	return DataClassificationEnum{
		FACTUAL_DATA: DataClassification{
			value: "FACTUAL_DATA",
		},
		DIMENSION_DATA: DataClassification{
			value: "DIMENSION_DATA",
		},
	}
}

func (c DataClassification) Value() string {
	return c.value
}

func (c DataClassification) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataClassification) UnmarshalJSON(b []byte) error {
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
