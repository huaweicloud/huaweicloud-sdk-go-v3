package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainConfigurationDto 修改域配置请求体。
type UpdateDomainConfigurationDto struct {

	// **参数说明**：自定义服务端证书ID
	ServerCertificateId *string `json:"server_certificate_id,omitempty"`

	ServerCertificateConfig *ServerCertificateConfig `json:"server_certificate_config,omitempty"`
}

func (o UpdateDomainConfigurationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainConfigurationDto struct{}"
	}

	return strings.Join([]string{"UpdateDomainConfigurationDto", string(data)}, " ")
}
