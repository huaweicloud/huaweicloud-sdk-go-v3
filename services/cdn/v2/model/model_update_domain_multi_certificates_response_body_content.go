package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDomainMultiCertificatesResponseBodyContent https配置。
type UpdateDomainMultiCertificatesResponseBodyContent struct {

	// **参数解释：** 域名列表 **约束限制：** 不涉及 **取值范围：** 不涉及
	DomainName string `json:"domain_name"`

	// https开关(0：\"关闭\"；1：\"设置证书\")。
	HttpsSwitch *int32 `json:"https_switch,omitempty"`

	// 回源方式:1：\"回源跟随\"；2：\"HTTP\"(默认)，3：https（自建）。
	AccessOriginWay *int32 `json:"access_origin_way,omitempty"`

	// 强制跳转HTTPS（0：不强制；1：强制） 。
	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`

	// http2.0（0：关闭；1：开启）
	Http2 *int32 `json:"http2,omitempty"`

	// **参数解释：** 证书名称 **约束限制：** 不涉及 **取值范围：** 不涉及
	CertName *string `json:"cert_name,omitempty"`

	// **参数解释：** 证书内容 **约束限制：** 不涉及 **取值范围：** 不涉及
	Certificate *string `json:"certificate,omitempty"`

	// 证书类型。取值意义： - 0：自有证书 - 2：SCM证书
	CertificateType *int32 `json:"certificate_type,omitempty"`

	// **参数解释：** 证书最近过期时间，Unix时间戳格式，单位为毫秒 **约束限制：** 不涉及 **取值范围：** 不涉及
	ExpirationTime *int64 `json:"expiration_time,omitempty"`

	// **参数解释：** 证书类型 **约束限制：** 不涉及 **取值范围：** - 0: 国际证书 - 1: 国密证书  **默认取值：** 不涉及
	CertificateCategory *int32 `json:"certificate_category,omitempty"`

	// **参数解释：** 双证书场景下的证书列表 **约束限制：** 不涉及
	Certificates *[]MultiCertificateBatchVo `json:"certificates,omitempty"`
}

func (o UpdateDomainMultiCertificatesResponseBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainMultiCertificatesResponseBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateDomainMultiCertificatesResponseBodyContent", string(data)}, " ")
}
