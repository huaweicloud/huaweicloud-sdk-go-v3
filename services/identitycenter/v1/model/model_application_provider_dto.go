package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationProviderDto 应用程序提供商
type ApplicationProviderDto struct {

	// 应用程序提供者URN
	ApplicationProviderUrn string `json:"application_provider_urn"`

	DisplayData *DisplayDataDto `json:"display_data,omitempty"`

	// 支持的协议
	FederationProtocol *ApplicationProviderDtoFederationProtocol `json:"federation_protocol,omitempty"`

	// 应用程序提供者唯一标识符（ID）
	ApplicationProviderId string `json:"application_provider_id"`
}

func (o ApplicationProviderDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationProviderDto struct{}"
	}

	return strings.Join([]string{"ApplicationProviderDto", string(data)}, " ")
}

type ApplicationProviderDtoFederationProtocol struct {
	value string
}

type ApplicationProviderDtoFederationProtocolEnum struct {
	SAML  ApplicationProviderDtoFederationProtocol
	OAUTH ApplicationProviderDtoFederationProtocol
}

func GetApplicationProviderDtoFederationProtocolEnum() ApplicationProviderDtoFederationProtocolEnum {
	return ApplicationProviderDtoFederationProtocolEnum{
		SAML: ApplicationProviderDtoFederationProtocol{
			value: "SAML",
		},
		OAUTH: ApplicationProviderDtoFederationProtocol{
			value: "OAUTH",
		},
	}
}

func (c ApplicationProviderDtoFederationProtocol) Value() string {
	return c.value
}

func (c ApplicationProviderDtoFederationProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationProviderDtoFederationProtocol) UnmarshalJSON(b []byte) error {
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
