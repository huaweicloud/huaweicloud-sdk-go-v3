package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateComponentConfigurationRequestBody struct {

	// API版本。
	ApiVersion string `json:"api_version"`

	// 资源种类。
	Kind CreateComponentConfigurationRequestBodyKind `json:"kind"`

	// 配置项列表。
	Items []ConfigurationItem `json:"items"`
}

func (o CreateComponentConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateComponentConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CreateComponentConfigurationRequestBody", string(data)}, " ")
}

type CreateComponentConfigurationRequestBodyKind struct {
	value string
}

type CreateComponentConfigurationRequestBodyKindEnum struct {
	CONFIGURATION CreateComponentConfigurationRequestBodyKind
}

func GetCreateComponentConfigurationRequestBodyKindEnum() CreateComponentConfigurationRequestBodyKindEnum {
	return CreateComponentConfigurationRequestBodyKindEnum{
		CONFIGURATION: CreateComponentConfigurationRequestBodyKind{
			value: "Configuration",
		},
	}
}

func (c CreateComponentConfigurationRequestBodyKind) Value() string {
	return c.value
}

func (c CreateComponentConfigurationRequestBodyKind) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateComponentConfigurationRequestBodyKind) UnmarshalJSON(b []byte) error {
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
