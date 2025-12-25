package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataLayering **参数解释**: 数据分层 - ODS 操作数据存储层 - DWS 数据汇总层 - ADS 应用数据服务层  **约束限制** 不涉及 **取值范围**: - ODS - DWS - ADS  **默认值** 不涉及
type DataLayering struct {
	value string
}

type DataLayeringEnum struct {
	ODS DataLayering
	DWS DataLayering
	ADS DataLayering
}

func GetDataLayeringEnum() DataLayeringEnum {
	return DataLayeringEnum{
		ODS: DataLayering{
			value: "ODS",
		},
		DWS: DataLayering{
			value: "DWS",
		},
		ADS: DataLayering{
			value: "ADS",
		},
	}
}

func (c DataLayering) Value() string {
	return c.value
}

func (c DataLayering) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataLayering) UnmarshalJSON(b []byte) error {
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
