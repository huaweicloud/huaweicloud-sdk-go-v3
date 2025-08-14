package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateSsoConfigurationReqBody description: the request body of UpdateSsoConfiguration
type UpdateSsoConfigurationReqBody struct {
	SsoConfiguration *SsoConfigurationDto `json:"sso_configuration"`

	// 配置类型
	ConfigurationType UpdateSsoConfigurationReqBodyConfigurationType `json:"configuration_type"`
}

func (o UpdateSsoConfigurationReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSsoConfigurationReqBody struct{}"
	}

	return strings.Join([]string{"UpdateSsoConfigurationReqBody", string(data)}, " ")
}

type UpdateSsoConfigurationReqBodyConfigurationType struct {
	value string
}

type UpdateSsoConfigurationReqBodyConfigurationTypeEnum struct {
	APP_AUTHENTICATION_CONFIGURATION UpdateSsoConfigurationReqBodyConfigurationType
	SESSION                          UpdateSsoConfigurationReqBodyConfigurationType
}

func GetUpdateSsoConfigurationReqBodyConfigurationTypeEnum() UpdateSsoConfigurationReqBodyConfigurationTypeEnum {
	return UpdateSsoConfigurationReqBodyConfigurationTypeEnum{
		APP_AUTHENTICATION_CONFIGURATION: UpdateSsoConfigurationReqBodyConfigurationType{
			value: "APP_AUTHENTICATION_CONFIGURATION",
		},
		SESSION: UpdateSsoConfigurationReqBodyConfigurationType{
			value: "SESSION",
		},
	}
}

func (c UpdateSsoConfigurationReqBodyConfigurationType) Value() string {
	return c.value
}

func (c UpdateSsoConfigurationReqBodyConfigurationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateSsoConfigurationReqBodyConfigurationType) UnmarshalJSON(b []byte) error {
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
