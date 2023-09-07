package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type Configuration struct {

	// 组件配置数据。
	Data *interface{} `json:"data,omitempty"`

	// 操作时间。
	OperatedAt *sdktime.SdkTime `json:"operated_at,omitempty"`

	// 操作ID。
	OperationId *string `json:"operation_id,omitempty"`

	// 组件配置类型，当前CAE支持组件配置如下11类。  - rds：云数据库RDS。  - cse：微服务引擎CSE。  - env：环境变量。  - access：访问方式。  - scaling：伸缩策略。  - volume：云存储配置。  - healthCheck：健康检查。  - lifecycle：生命周期管理。  - apm2：性能管理。  - log: 自定义日志路径。  - customMetric: 自定义监控指标。
	Type *ConfigurationType `json:"type,omitempty"`

	// 配置是否生效。
	IsActivated *bool `json:"is_activated,omitempty"`
}

func (o Configuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configuration struct{}"
	}

	return strings.Join([]string{"Configuration", string(data)}, " ")
}

type ConfigurationType struct {
	value string
}

type ConfigurationTypeEnum struct {
	RDS           ConfigurationType
	CSE           ConfigurationType
	ENV           ConfigurationType
	ACCESS        ConfigurationType
	SCALING       ConfigurationType
	VOLUME        ConfigurationType
	HEALTH_CHECK  ConfigurationType
	LIFECYCLE     ConfigurationType
	APM2          ConfigurationType
	LOG           ConfigurationType
	CUSTOM_METRIC ConfigurationType
}

func GetConfigurationTypeEnum() ConfigurationTypeEnum {
	return ConfigurationTypeEnum{
		RDS: ConfigurationType{
			value: "rds",
		},
		CSE: ConfigurationType{
			value: "cse",
		},
		ENV: ConfigurationType{
			value: "env",
		},
		ACCESS: ConfigurationType{
			value: "access",
		},
		SCALING: ConfigurationType{
			value: "scaling",
		},
		VOLUME: ConfigurationType{
			value: "volume",
		},
		HEALTH_CHECK: ConfigurationType{
			value: "healthCheck",
		},
		LIFECYCLE: ConfigurationType{
			value: "lifecycle",
		},
		APM2: ConfigurationType{
			value: "apm2",
		},
		LOG: ConfigurationType{
			value: "log",
		},
		CUSTOM_METRIC: ConfigurationType{
			value: "customMetric",
		},
	}
}

func (c ConfigurationType) Value() string {
	return c.value
}

func (c ConfigurationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationType) UnmarshalJSON(b []byte) error {
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
