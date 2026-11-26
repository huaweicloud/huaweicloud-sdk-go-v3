package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainMultiCertificatesRequestBodyContent **参数解释：** HTTPS证书配置，配置证书后，允许客户端采用HTTPS协议访问CDN节点 **约束限制：** 不涉及
type UpdateDomainMultiCertificatesRequestBodyContent struct {

	// **参数解释：** 需要配置HTTPS证书的域名 **约束限制：** - 域名间以“,”分隔 - 最多配置50个域名  **取值范围：** 不涉及 **默认取值：** 不涉及
	DomainName string `json:"domain_name"`

	// **参数解释：** 证书配置开关 **约束限制：** https_switch为1时，证书参数不能为空 **取值范围：** - 0: 关闭 - 1: 开启  **默认取值：** 不涉及
	HttpsSwitch int32 `json:"https_switch"`

	// **参数解释：** 回源协议 **约束限制：** 不涉及 **取值范围：** - 1: 协议跟随 - 2: http协议 - 3: https协议  **默认取值：** 2: http协议
	AccessOriginWay *int32 `json:"access_origin_way,omitempty"`

	// **参数解释：** 是否开启HTTPS强制跳转，强制客户端使用HTTPS协议访问CDN节点 > 此参数即将下线，建议使用force_redirect_config修改配置  **约束限制：** 不涉及 **取值范围：** - 0: 关闭 - 1: 开启  **默认取值：** 0: 关闭
	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	// **参数解释：** HTTP/2开关，支持客户端采用HTTP/2协议访问CDN节点 **约束限制：** 不涉及 **取值范围：** - 0: 关闭 - 1: 开启  **默认取值：** 0: 关闭
	Http2 *int32 `json:"http2,omitempty"`

	// **参数解释：** 证书名称 **约束限制：** https_switch为1时，该参数必填 **取值范围：** 3-64个字符，由英文字母、数字及“-”、“_”、空格组成 **默认取值：** 不涉及
	CertName *string `json:"cert_name,omitempty"`

	// **参数解释：** SSL证书内容 **约束限制：** - 仅支持PEM编码格式 - 不启用证书则无需输入 - 初次配置证书时必传 - 需要传入完整证书链  **取值范围：** 证书链长度不超过20KB **默认取值：** 不涉及
	Certificate *string `json:"certificate,omitempty"`

	// **参数解释：** SSL证书私钥内容 **约束限制：** - 仅支持PEM编码格式 - 不启用证书则无需输入 - 初次配置证书时必传  **取值范围：** 不涉及 **默认取值：** 不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// **参数解释：** 证书类型 **约束限制：** 不涉及 **取值范围：** - 0: 自有证书 - 2: SCM证书  **默认取值：** 0: 自有证书
	CertificateType *int32 `json:"certificate_type,omitempty"`

	// **参数解释：** SCM证书ID，可使用云证书与管理服务（CCM）的查询证书列表接口获取 **约束限制：** 配置SCM证书时必传 **取值范围：** 不涉及 **默认取值：** 不涉及
	ScmCertificateId *string `json:"scm_certificate_id,omitempty"`

	// **参数解释：** 配置国际证书和国密证书双证书场景下的证书列表 **约束限制：** - 配置双证书时最多传2个证书，只能包含1个国际证书和1个国密证书 - 使用双证书配置时，cert_name/certificate/private_key等单证书参数不可再使用  **取值范围：** 不涉及
	Certificates *[]MultiCertificateBatchVo `json:"certificates,omitempty"`
}

func (o UpdateDomainMultiCertificatesRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesRequestBodyContent", string(data)}, " ")
}
