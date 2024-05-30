package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModelLevel 数据治理分层。 枚举值：   - SDI: 贴源数据层   - DWI: 数据整合层   - DWR: 数据报告层   - DM:  数据集市层
type ModelLevel struct {
	value string
}

type ModelLevelEnum struct {
	SDI ModelLevel
	DWI ModelLevel
	DWR ModelLevel
	DM  ModelLevel
}

func GetModelLevelEnum() ModelLevelEnum {
	return ModelLevelEnum{
		SDI: ModelLevel{
			value: "SDI",
		},
		DWI: ModelLevel{
			value: "DWI",
		},
		DWR: ModelLevel{
			value: "DWR",
		},
		DM: ModelLevel{
			value: "DM",
		},
	}
}

func (c ModelLevel) Value() string {
	return c.value
}

func (c ModelLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModelLevel) UnmarshalJSON(b []byte) error {
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
