package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainConfigurationDto 域配置结构体
type DomainConfigurationDto struct {

	// 域配置唯一标识ID
	ConfigurationId *string `json:"configuration_id,omitempty"`

	// **参数说明**：自定义域名
	DomainName *string `json:"domain_name,omitempty"`

	// **参数说明**：应用协议场景，当前仅支持Device-MQTTS：设备接入MQTTS场景
	AccessProtocol *string `json:"access_protocol,omitempty"`

	// **参数说明**：自定义服务端证书ID
	ServerCertificateId *string `json:"server_certificate_id,omitempty"`

	ServerCertificateConfig *ServerCertificateConfig `json:"server_certificate_config,omitempty"`
}

func (o DomainConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainConfigurationDto struct{}"
	}

	return strings.Join([]string{"DomainConfigurationDto", string(data)}, " ")
}
