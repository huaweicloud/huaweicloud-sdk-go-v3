package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApplicationInstanceDto 应用程序实例
type ApplicationInstanceDto struct {
	ActiveCertificate *CertificateDto `json:"active_certificate"`

	Display *DisplayDto `json:"display"`

	IdentityProviderConfig *IdentityProviderConfigDto `json:"identity_provider_config"`

	// 应用程序实例唯一标识ID
	ApplicationInstanceId string `json:"application_instance_id"`

	// 应用程序UUID
	Name string `json:"name"`

	// 应用程序在门户是否可见
	Visible bool `json:"visible"`

	ResponseConfig *ResponseConfigDto `json:"response_config"`

	ResponseSchemaConfig *ResponseSchemaConfigDto `json:"response_schema_config"`

	SecurityConfig *SecurityConfigDto `json:"security_config"`

	// 应用程序实例状态
	Status ApplicationInstanceDtoStatus `json:"status"`

	Template *ApplicationTemplateDto `json:"template"`

	ServiceProviderConfig *ServiceProviderConfigDto `json:"service_provider_config"`

	// OIDC Client ID
	ClientId *string `json:"client_id,omitempty"`

	// 用户是否可见
	EndUserVisible *bool `json:"end_user_visible,omitempty"`

	// 组员所属账号ID
	ManagedAccount *string `json:"managed_account,omitempty"`
}

func (o ApplicationInstanceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationInstanceDto struct{}"
	}

	return strings.Join([]string{"ApplicationInstanceDto", string(data)}, " ")
}

type ApplicationInstanceDtoStatus struct {
	value string
}

type ApplicationInstanceDtoStatusEnum struct {
	CREATED ApplicationInstanceDtoStatus
	ENABLED ApplicationInstanceDtoStatus
}

func GetApplicationInstanceDtoStatusEnum() ApplicationInstanceDtoStatusEnum {
	return ApplicationInstanceDtoStatusEnum{
		CREATED: ApplicationInstanceDtoStatus{
			value: "CREATED",
		},
		ENABLED: ApplicationInstanceDtoStatus{
			value: "ENABLED",
		},
	}
}

func (c ApplicationInstanceDtoStatus) Value() string {
	return c.value
}

func (c ApplicationInstanceDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApplicationInstanceDtoStatus) UnmarshalJSON(b []byte) error {
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
