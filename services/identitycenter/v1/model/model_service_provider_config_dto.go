package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceProviderConfigDto 服务提供商
type ServiceProviderConfigDto struct {

	// SAML受众
	Audience string `json:"audience"`

	// 是否需要签名
	RequireRequestSignature *bool `json:"require_request_signature,omitempty"`

	// SAML响应接收方
	Consumers []ConsumersDto `json:"consumers"`

	// 应用程序启动Url
	StartUrl *string `json:"start_url,omitempty"`
}

func (o ServiceProviderConfigDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceProviderConfigDto struct{}"
	}

	return strings.Join([]string{"ServiceProviderConfigDto", string(data)}, " ")
}
