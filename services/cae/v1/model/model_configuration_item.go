package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConfigurationItem struct {

	// 组件配置类型。
	Type ConfigurationItemType `json:"type"`

	Data *ConfigurationData `json:"data"`
}

func (o ConfigurationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationItem struct{}"
	}

	return strings.Join([]string{"ConfigurationItem", string(data)}, " ")
}

type ConfigurationItemType struct {
	value string
}

type ConfigurationItemTypeEnum struct {
	RDS           ConfigurationItemType
	CSE           ConfigurationItemType
	ENV           ConfigurationItemType
	ACCESS        ConfigurationItemType
	SCALING       ConfigurationItemType
	VOLUME        ConfigurationItemType
	HEALTH_CHECK  ConfigurationItemType
	LIFECYCLE     ConfigurationItemType
	APM2          ConfigurationItemType
	LOG           ConfigurationItemType
	CUSTOM_METRIC ConfigurationItemType
}

func GetConfigurationItemTypeEnum() ConfigurationItemTypeEnum {
	return ConfigurationItemTypeEnum{
		RDS: ConfigurationItemType{
			value: "rds",
		},
		CSE: ConfigurationItemType{
			value: "cse",
		},
		ENV: ConfigurationItemType{
			value: "env",
		},
		ACCESS: ConfigurationItemType{
			value: "access",
		},
		SCALING: ConfigurationItemType{
			value: "scaling",
		},
		VOLUME: ConfigurationItemType{
			value: "volume",
		},
		HEALTH_CHECK: ConfigurationItemType{
			value: "healthCheck",
		},
		LIFECYCLE: ConfigurationItemType{
			value: "lifecycle",
		},
		APM2: ConfigurationItemType{
			value: "apm2",
		},
		LOG: ConfigurationItemType{
			value: "log",
		},
		CUSTOM_METRIC: ConfigurationItemType{
			value: "customMetric",
		},
	}
}

func (c ConfigurationItemType) Value() string {
	return c.value
}

func (c ConfigurationItemType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationItemType) UnmarshalJSON(b []byte) error {
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
