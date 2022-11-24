package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LanguageEnum struct {
	value string
}

type LanguageEnumEnum struct {
	EN_US LanguageEnum
	ZH_CN LanguageEnum
}

func GetLanguageEnumEnum() LanguageEnumEnum {
	return LanguageEnumEnum{
		EN_US: LanguageEnum{
			value: "en_us",
		},
		ZH_CN: LanguageEnum{
			value: "zh_cn",
		},
	}
}

func (c LanguageEnum) Value() string {
	return c.value
}

func (c LanguageEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LanguageEnum) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
