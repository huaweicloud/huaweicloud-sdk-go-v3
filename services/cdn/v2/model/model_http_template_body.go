package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HttpTemplateBody **参数解释：** 在CDN节点添加证书，支持客户端以HTTPS协议访问节点，实现HTTPS安全加速 **约束限制：** 不涉及
type HttpTemplateBody struct {

	// **参数解释：** 证书配置开关 **约束限制：** 不涉及 **取值范围：** - on: 开启 - off: 关闭 **默认取值：** 不涉及
	HttpsStatus *string `json:"https_status,omitempty"`

	// **参数解释：** 证书类型 **约束限制：** 不涉及 **取值范围：** - server: 国际证书 - server_sm: 国密证书 **默认取值：** 不涉及
	CertificateType *string `json:"certificate_type,omitempty"`

	// **参数解释：** 证书来源 **约束限制：** 证书开启时必传 **取值范围：** - 0: 自有证书 - 2: SCM证书 **默认取值：** 不涉及
	CertificateSource *int32 `json:"certificate_source,omitempty"`

	// **参数解释：** SCM证书ID，可使用云证书与管理服务（CCM）的查询证书列表接口获取 **约束限制：** 配置SCM证书时必传 **取值范围：** 不涉及 **默认取值：** 不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// **参数解释：** 证书名称 **约束限制：** - 当证书开启时必传 - 3-64个字符 **取值范围：** 不涉及 **默认取值：** 不涉及
	CertificateName *string `json:"certificate_name,omitempty"`

	// **参数解释：** HTTPS协议使用的证书内容 **约束限制：** - 仅支持PEM编码格式 - 当证书开启时必传 -  需要传入完整证书链 **取值范围：** 证书链长度不超过20KB **默认取值：** 不涉及
	CertificateValue *string `json:"certificate_value,omitempty"`

	// **参数解释：** HTTPS协议使用的私钥 **约束限制：** - 当证书开启时必传 - 仅支持PEM编码格式 **取值范围：** 不涉及 **默认取值：** 不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// **参数解释：** 加密证书内容 **约束限制：** - 证书类型为国密证书时必传 - 仅支持PEM编码格式 **取值范围：** 不涉及 **默认取值：** 不涉及
	EncCertificateValue *string `json:"enc_certificate_value,omitempty"`

	// **参数解释：** 加密私钥内容 **约束限制：** - 证书类型为国密证书时必传 - 仅支持PEM编码格式 **取值范围：** 不涉及 **默认取值：** 不涉及
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	Certificates *[]CertificatesPutBody `json:"certificates,omitempty"`

	// **参数解释：** HTTP/2开关，支持客户端采用HTTP/2协议访问CDN节点 **约束限制：** https_status=off时，该值不生效 **取值范围：** - on: 是 - off: 否 **默认取值：** off: 否
	Http2Status *string `json:"http2_status,omitempty"`

	// **参数解释：** 是否开启ocsp stapling **约束限制：** 不涉及 **取值范围：** - on: 是 - off: 否 **默认取值：** 不涉及
	OcspStaplingStatus *string `json:"ocsp_stapling_status,omitempty"`
}

func (o HttpTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpTemplateBody struct{}"
	}

	return strings.Join([]string{"HttpTemplateBody", string(data)}, " ")
}
