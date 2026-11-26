package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClientCert **参数解释：** 客户端证书配置，实现客户端与CDN节点的双向证书认证，提高网站通信安全性 **约束限制：** 不涉及
type ClientCert struct {

	// **参数解释：** 客户端证书配置开关 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭  **默认取值：** 不涉及
	Status string `json:"status"`

	// **参数解释：** 客户端CA证书的内容 **约束限制：** 仅支持PEM格式 **取值范围：** 不涉及 **默认取值：** 不涉及
	TrustedCert *string `json:"trusted_cert,omitempty"`

	// **参数解释：** 客户端CA证书指定的域名 **约束限制：** - 如果不配置域名，则CDN会放行所有持有该CA证书的客户端请求 - 最多可配置100个域名 - 多个域名用“,”或“|”分隔  **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *string `json:"hosts,omitempty"`

	// **参数解释：** 多证书场景下客户端CA证书列表 **约束限制：** 不涉及
	TrustedCerts *[]ClientCertTrustedCerts `json:"trusted_certs,omitempty"`

	// **参数解释：** 国密证书场景下客户端CA证书列表 **约束限制：** - 开启国密证书后配置 - 仅支持PEM格式
	SmTrustedCerts *[]ClientCertTrustedCerts `json:"sm_trusted_certs,omitempty"`

	// **参数解释：** 国密证书指定访问的客户端域名 **约束限制：** - 开启国密证书后配置 - 最多可配置100个域名 - 多个域名用“|”分隔  **取值范围：** 不涉及 **默认取值：** 不涉及
	SmCertHosts *string `json:"sm_cert_hosts,omitempty"`
}

func (o ClientCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientCert struct{}"
	}

	return strings.Join([]string{"ClientCert", string(data)}, " ")
}
