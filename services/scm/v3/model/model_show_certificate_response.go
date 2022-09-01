package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCertificateResponse struct {

	// 证书id。
	Id *string `json:"id,omitempty" xml:"id"`

	// 证书状态。取值如下： - PAID：证书已支付，待申请证书。 - ISSUED：证书已签发。 - CHECKING：证书申请审核中。 - CANCELCHECKING：取消证书申请审核中。 - UNPASSED：证书申请未通过。 - EXPIRED：证书已过期。 - REVOKING：证书吊销申请审核中。 - REVOKED：证书已吊销。 - UPLOAD：证书托管中。 - SUPPLEMENTCHECKING：多域名证书新增附加域名审核中。 - CANCELSUPPLEMENTING：取消新增附加域名审核中。
	Status *string `json:"status,omitempty" xml:"status"`

	// 订单id。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 证书名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 证书类型。取值如下： DV_SSL_CERT、DV_SSL_CERT_BASIC、EV_SSL_CERT、 EV_SSL_CERT_PRO、OV_SSL_CERT、OV_SSL_CERT_PRO。
	Type *string `json:"type,omitempty" xml:"type"`

	// 证书品牌。取值如下： GLOBALSIGN、SYMANTEC、GEOTRUST、CFCA。
	Brand *string `json:"brand,omitempty" xml:"brand"`

	// 证书是否支持推送。
	PushSupport *string `json:"push_support,omitempty" xml:"push_support"`

	// 证书吊销原因。
	RevokeReason *string `json:"revoke_reason,omitempty" xml:"revoke_reason"`

	// 签名算法。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty" xml:"signature_algorithm"`

	// 证书签发时间，没有获取到有效值时为空。
	IssueTime *string `json:"issue_time,omitempty" xml:"issue_time"`

	// 证书生效时间，没有获取到有效值时为空。
	NotBefore *string `json:"not_before,omitempty" xml:"not_before"`

	// 证书失效时间，没有获取到有效值时为空。
	NotAfter *string `json:"not_after,omitempty" xml:"not_after"`

	// 证书有效期，按月为单位。
	ValidityPeriod *int32 `json:"validity_period,omitempty" xml:"validity_period"`

	// 域名认证方式，取值如下：DNS、FILE、EMAIL。
	ValidationMethod *string `json:"validation_method,omitempty" xml:"validation_method"`

	// 域名类型，取值如下： - SINGLE_DOMAIN：单域名 - WILDCARD：通配符 - MULTI_DOMAIN：多域名
	DomainType *string `json:"domain_type,omitempty" xml:"domain_type"`

	// 证书绑定域名。
	Domain *string `json:"domain,omitempty" xml:"domain"`

	// 证书绑定的附加域名信息。
	Sans *string `json:"sans,omitempty" xml:"sans"`

	// 证书可绑定域名个数。
	DomainCount *int32 `json:"domain_count,omitempty" xml:"domain_count"`

	// 证书可绑定附加域名个数。
	WildcardCount *int32 `json:"wildcard_count,omitempty" xml:"wildcard_count"`

	// 企业项目ID，默认为“0”。 对于开通企业项目的用户，表示资源处于默认企业项目下。 对于未开通企业项目的用户，表示资源未处于企业项目下。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 域名所有权认证信息，详情请参见Authentification字段数据结构说明。
	Authentification *[]Authentification `json:"authentification,omitempty" xml:"authentification"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ShowCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateResponse", string(data)}, " ")
}
