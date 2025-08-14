package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateCertificateDto 更新CA证书结构体。
type UpdateCertificateDto struct {

	// 是否开启自注册能力，当为true时该功能必须配合预调配功能使用，true：是，false：否。
	ProvisionEnable *bool `json:"provision_enable,omitempty"`

	// 预调配模板ID，该CA证书绑定的预调配模板id，当该字段传null时表示解除绑定关系。
	TemplateId *string `json:"template_id,omitempty"`

	// 是否开启该CA证书下的设备证书OCSP校验，当为true且设备证书信息中包含OCSP url时则平台会校验证书的状态，当证书状态为revoked时平台会拒绝设备连接，true：开启，false：关闭。
	OcspEnable *bool `json:"ocsp_enable,omitempty"`

	// 访问ocsp服务器是否开启SSL，开启后必须配置服务器CA证书校验。
	OcspSslEnable *bool `json:"ocsp_ssl_enable,omitempty"`

	// ocsp服务器端CA证书id，当ocsp服务器为https协议时需要配置，否则认证失败。
	OcspServerCaId *string `json:"ocsp_server_ca_id,omitempty"`
}

func (o UpdateCertificateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateDto struct{}"
	}

	return strings.Join([]string{"UpdateCertificateDto", string(data)}, " ")
}
