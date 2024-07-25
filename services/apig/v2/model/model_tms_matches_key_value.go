package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TmsMatchesKeyValue struct {

	// 要匹配的字段 - resource_name 表示按照APIG实例的名称去匹配
	Key *TmsMatchesKeyValueKey `json:"key,omitempty"`

	// 值。 支持可用 UTF-8 格式表示的字母(包含中文)、数字和空格，以及以下字符： _ . : / = + - @
	Value *string `json:"value,omitempty"`
}

func (o TmsMatchesKeyValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsMatchesKeyValue struct{}"
	}

	return strings.Join([]string{"TmsMatchesKeyValue", string(data)}, " ")
}

type TmsMatchesKeyValueKey struct {
	value string
}

type TmsMatchesKeyValueKeyEnum struct {
	RESOURCE_NAME TmsMatchesKeyValueKey
}

func GetTmsMatchesKeyValueKeyEnum() TmsMatchesKeyValueKeyEnum {
	return TmsMatchesKeyValueKeyEnum{
		RESOURCE_NAME: TmsMatchesKeyValueKey{
			value: "resource_name",
		},
	}
}

func (c TmsMatchesKeyValueKey) Value() string {
	return c.value
}

func (c TmsMatchesKeyValueKey) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TmsMatchesKeyValueKey) UnmarshalJSON(b []byte) error {
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
