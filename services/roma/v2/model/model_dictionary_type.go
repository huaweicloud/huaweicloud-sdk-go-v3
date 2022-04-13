package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 字典类型 - system: 代表系统内置字典,用户可编辑但不可删除 - user: 代表用户创建字典
type DictionaryType struct {
	value string
}

type DictionaryTypeEnum struct {
	SYSTEM DictionaryType
	USER   DictionaryType
}

func GetDictionaryTypeEnum() DictionaryTypeEnum {
	return DictionaryTypeEnum{
		SYSTEM: DictionaryType{
			value: "system",
		},
		USER: DictionaryType{
			value: "user",
		},
	}
}

func (c DictionaryType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DictionaryType) UnmarshalJSON(b []byte) error {
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
