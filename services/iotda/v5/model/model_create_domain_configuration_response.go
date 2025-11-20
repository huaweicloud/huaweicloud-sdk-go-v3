package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainConfigurationResponse Response Object
type CreateDomainConfigurationResponse struct {

	// 域配置唯一标识ID
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// **参数说明**：自定义域名
	DomainName *string `json:"domain_name,omitempty"`

	// **参数说明**：应用协议场景，当前仅支持Device-MQTTS：设备接入MQTTS场景
	AccessProtocol *string `json:"access_protocol,omitempty"`

	// **参数说明**：自定义服务端证书ID
	ServerCertificateId *string `json:"server_certificate_id,omitempty"`

	ServerCertificateConfig *ServerCertificateConfig `json:"server_certificate_config,omitempty"`
	HttpStatusCode          int                      `json:"-"`
}

func (o CreateDomainConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainConfigurationResponse struct{}"
	}

	return strings.Join([]string{"CreateDomainConfigurationResponse", string(data)}, " ")
}
