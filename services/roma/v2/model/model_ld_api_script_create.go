package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LdApiScriptCreate struct {

	// API类型 - data：数据API - function：函数API
	ApiType *LdApiScriptCreateApiType `json:"api_type,omitempty"`

	// API脚本信息列表
	Scripts *[]LdApiScriptBase `json:"scripts,omitempty"`
}

func (o LdApiScriptCreate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LdApiScriptCreate struct{}"
	}

	return strings.Join([]string{"LdApiScriptCreate", string(data)}, " ")
}

type LdApiScriptCreateApiType struct {
	value string
}

type LdApiScriptCreateApiTypeEnum struct {
	DATA     LdApiScriptCreateApiType
	FUNCTION LdApiScriptCreateApiType
}

func GetLdApiScriptCreateApiTypeEnum() LdApiScriptCreateApiTypeEnum {
	return LdApiScriptCreateApiTypeEnum{
		DATA: LdApiScriptCreateApiType{
			value: "data",
		},
		FUNCTION: LdApiScriptCreateApiType{
			value: "function",
		},
	}
}

func (c LdApiScriptCreateApiType) Value() string {
	return c.value
}

func (c LdApiScriptCreateApiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LdApiScriptCreateApiType) UnmarshalJSON(b []byte) error {
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
