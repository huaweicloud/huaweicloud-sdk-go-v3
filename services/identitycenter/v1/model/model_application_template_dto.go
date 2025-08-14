package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationTemplateDto 应用程序模板
type ApplicationTemplateDto struct {
	Application *ApplicationTemplateDisplayDto `json:"application"`

	ResponseConfig *ResponseConfigDto `json:"response_config"`

	ResponseSchemaConfig *ResponseSchemaConfigDto `json:"response_schema_config"`

	// 支持的协议
	SsoProtocol string `json:"sso_protocol"`

	SecurityConfig *SecurityConfigDto `json:"security_config"`

	ServiceProviderConfig *ServiceProviderConfigDto `json:"service_provider_config"`

	// 应用程序模板唯一标识ID
	TemplateId string `json:"template_id"`

	// 应用程序模板版本
	TemplateVersion string `json:"template_version"`
}

func (o ApplicationTemplateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationTemplateDto struct{}"
	}

	return strings.Join([]string{"ApplicationTemplateDto", string(data)}, " ")
}
