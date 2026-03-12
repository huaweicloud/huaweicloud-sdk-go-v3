package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConfigurationParameterResult struct {

	// 参数名称。
	Name *string `json:"name,omitempty"`

	// 参数值。
	Value *string `json:"value,omitempty"`

	// 修改该参数是否需要重启实例。
	RestartRequired *bool `json:"restart_required,omitempty"`

	// **参数解释**: 该参数的value值是否为只读，无法直接修改。 **取值范围**: - true：该参数的value值只读，不允许用户直接修改。 - false：允许修改。
	Readonly *bool `json:"readonly,omitempty"`

	// 参数取值范围。
	ValueRange *string `json:"value_range,omitempty"`

	// 参数类型，取值为“string”、“integer”、“boolean”、“list”或“float”之一。
	Type *ConfigurationParameterResultType `json:"type,omitempty"`

	// 参数描述。
	Description *string `json:"description,omitempty"`
}

func (o ConfigurationParameterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationParameterResult struct{}"
	}

	return strings.Join([]string{"ConfigurationParameterResult", string(data)}, " ")
}

type ConfigurationParameterResultType struct {
	value string
}

type ConfigurationParameterResultTypeEnum struct {
	STRING  ConfigurationParameterResultType
	INTEGER ConfigurationParameterResultType
	BOOLEAN ConfigurationParameterResultType
	LIST    ConfigurationParameterResultType
	FLOAT   ConfigurationParameterResultType
}

func GetConfigurationParameterResultTypeEnum() ConfigurationParameterResultTypeEnum {
	return ConfigurationParameterResultTypeEnum{
		STRING: ConfigurationParameterResultType{
			value: "string",
		},
		INTEGER: ConfigurationParameterResultType{
			value: "integer",
		},
		BOOLEAN: ConfigurationParameterResultType{
			value: "boolean",
		},
		LIST: ConfigurationParameterResultType{
			value: "list",
		},
		FLOAT: ConfigurationParameterResultType{
			value: "float",
		},
	}
}

func (c ConfigurationParameterResultType) Value() string {
	return c.value
}

func (c ConfigurationParameterResultType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationParameterResultType) UnmarshalJSON(b []byte) error {
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
