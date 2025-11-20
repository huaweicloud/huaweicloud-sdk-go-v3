package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDomainConfigurationDto 创建域配置请求体。
type CreateDomainConfigurationDto struct {

	// **参数说明**：自定义域名
	DomainName string `json:"domain_name"`

	// **参数说明**：接入协议，当前仅支持Device-MQTTS：设备接入MQTTS场景
	AccessProtocol string `json:"access_protocol"`

	// **参数说明**：自定义服务端证书ID
	ServerCertificateId string `json:"server_certificate_id"`

	ServerCertificateConfig *ServerCertificateConfig `json:"server_certificate_config,omitempty"`
}

func (o CreateDomainConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainConfigurationDto struct{}"
	}

	return strings.Join([]string{"CreateDomainConfigurationDto", string(data)}, " ")
}
