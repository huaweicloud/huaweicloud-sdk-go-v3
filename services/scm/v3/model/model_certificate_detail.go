package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateDetail struct {
	// 证书id。

	Id *string `json:"id,omitempty"`
	// 证书名称。

	Name *string `json:"name,omitempty"`
	// 证书绑定的域名。

	Domain *string `json:"domain,omitempty"`
	// 证书类型。取值如下： DV_SSL_CERT、DV_SSL_CERT_BASIC、EV_SSL_CERT、 EV_SSL_CERT_PRO、OV_SSL_CERT、OV_SSL_CERT_PRO

	Type *string `json:"type,omitempty"`
	// 证书品牌。取值如下：GLOBALSIGN、SYMANTEC、GEOTRUST、CFCA

	Brand *string `json:"brand,omitempty"`
	// 证书过期时间。

	ExpireTime *string `json:"expire_time,omitempty"`
	// 域名类型。取值如下： - SINGLE_DOMAIN：单域名 - WILDCARD：通配符 - MULTI_DOMAIN：多域名

	DomainType *string `json:"domain_type,omitempty"`
	// 证书有效期，以月为单位。

	ValidityPeriod *int32 `json:"validity_period,omitempty"`
	// 证书状态，取值如下： - PAID：证书已支付；待申请证书。 - ISSUED：证书已签发。 - CHECKING：证书申请审核中。 - CANCELCHECKING：取消证书申请审核中。 - UNPASSED：证书申请未通过。 - EXPIRED：证书已过期。 - REVOKING：证书吊销申请审核中。 - CANCLEREVOKING：证书取消吊销申请审核中。 - REVOKED：证书已吊销。 - UPLOAD：证书托管中。 - SUPPLEMENTCHECKING：多域名证书新增附加域名审核中。 - CANCELSUPPLEMENTING：取消新增附加域名审核中。

	Status *string `json:"status,omitempty"`
	// 证书可绑定域名个数。

	DomainCount *int32 `json:"domain_count,omitempty"`
	// 证书可绑定泛域名个数。

	WildcardCount *int32 `json:"wildcard_count,omitempty"`
	// 证书描述。

	Description *string `json:"description,omitempty"`
}

func (o CertificateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateDetail struct{}"
	}

	return strings.Join([]string{"CertificateDetail", string(data)}, " ")
}
