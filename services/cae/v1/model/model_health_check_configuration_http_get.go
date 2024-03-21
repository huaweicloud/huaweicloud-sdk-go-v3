package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// HealthCheckConfigurationHttpGet 健康检查配置中HTTP请求检查信息。
type HealthCheckConfigurationHttpGet struct {

	// URL路径。
	Path *string `json:"path,omitempty"`

	// 端口。
	Port *int32 `json:"port,omitempty"`

	// 协议。
	Scheme *HealthCheckConfigurationHttpGetScheme `json:"scheme,omitempty"`
}

func (o HealthCheckConfigurationHttpGet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HealthCheckConfigurationHttpGet struct{}"
	}

	return strings.Join([]string{"HealthCheckConfigurationHttpGet", string(data)}, " ")
}

type HealthCheckConfigurationHttpGetScheme struct {
	value string
}

type HealthCheckConfigurationHttpGetSchemeEnum struct {
	HTTP  HealthCheckConfigurationHttpGetScheme
	HTTPS HealthCheckConfigurationHttpGetScheme
}

func GetHealthCheckConfigurationHttpGetSchemeEnum() HealthCheckConfigurationHttpGetSchemeEnum {
	return HealthCheckConfigurationHttpGetSchemeEnum{
		HTTP: HealthCheckConfigurationHttpGetScheme{
			value: "HTTP",
		},
		HTTPS: HealthCheckConfigurationHttpGetScheme{
			value: "HTTPS",
		},
	}
}

func (c HealthCheckConfigurationHttpGetScheme) Value() string {
	return c.value
}

func (c HealthCheckConfigurationHttpGetScheme) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HealthCheckConfigurationHttpGetScheme) UnmarshalJSON(b []byte) error {
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
