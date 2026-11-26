package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HttpsDetail struct {

	// **参数解释：** 域名id。 **约束限制：** 不涉及 **取值范围：** 不涉及
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 绑定该证书的域名。 **约束限制：** 不涉及 **取值范围：** 不涉及
	DomainName *string `json:"domain_name,omitempty"`

	// **参数解释：** 证书名字。 **约束限制：** 不涉及 **取值范围：** 不涉及
	CertName *string `json:"cert_name,omitempty"`

	// **参数解释：** 证书内容。 **约束限制：** 不涉及 **取值范围：** 不涉及
	Certificate *string `json:"certificate,omitempty"`

	// **参数解释：** 私钥内容。 **约束限制：** 不涉及 **取值范围：** 不涉及
	PrivateKey *string `json:"private_key,omitempty"`

	// 证书类型，0：自有证书，1：云托管证书，2：SCM证书。
	CertificateType *int32 `json:"certificate_type,omitempty"`

	// **参数解释：** 证书过期时间，Unix时间戳格式，单位为毫秒。 **约束限制：** 不涉及 **取值范围：** 不涉及
	ExpirationTime *int64 `json:"expiration_time,omitempty"`

	// HTTPS证书是否启用，取值0：不启用；1：启用HTTPS加速并协议跟随回源；2：启用HTTPS加速并HTTP回源。
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// 客户端请求是否强制重定向。1是，0否。（如果为2，表示强制跳转HTTP）
	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	// 是否使用HTTP2.0。（1是，0否。）
	Http2 *int32 `json:"http2,omitempty"`

	// **参数解释：** 证书列表，配置双证书场景下包含国际证书和国密证书两项 **约束限制：** 不涉及
	Certificates *[]HttpsDetailCertificates `json:"certificates,omitempty"`
}

func (o HttpsDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HttpsDetail struct{}"
	}

	return strings.Join([]string{"HttpsDetail", string(data)}, " ")
}
