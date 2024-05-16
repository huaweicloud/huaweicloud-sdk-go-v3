package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateResponse Response Object
type ShowCertificateResponse struct {

	// 证书id。
	Id *string `json:"id,omitempty"`

	// 证书状态。取值如下： - PAID：证书已支付，待申请证书。 - ISSUED：证书已签发。 - CHECKING：证书申请审核中。 - CANCELCHECKING：取消证书申请审核中。 - UNPASSED：证书申请未通过。 - EXPIRED：证书已过期。 - REVOKING：证书吊销申请审核中。 - REVOKED：证书已吊销。 - UPLOAD：证书托管中。 - SUPPLEMENTCHECKING：多域名证书新增附加域名审核中。 - CANCELSUPPLEMENTING：取消新增附加域名审核中。
	Status *string `json:"status,omitempty"`

	// 订单id。
	OrderId *string `json:"order_id,omitempty"`

	// 证书名称。
	Name *string `json:"name,omitempty"`

	// 证书类型。取值如下： DV_SSL_CERT、DV_SSL_CERT_BASIC、EV_SSL_CERT、 EV_SSL_CERT_PRO、OV_SSL_CERT、OV_SSL_CERT_PRO。
	Type *string `json:"type,omitempty"`

	// 证书品牌。取值如下： GLOBALSIGN、SYMANTEC、GEOTRUST、CFCA。
	Brand *string `json:"brand,omitempty"`

	// 证书是否支持推送。
	PushSupport *string `json:"push_support,omitempty"`

	// 证书吊销原因。
	RevokeReason *string `json:"revoke_reason,omitempty"`

	// 签名算法。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 证书签发时间，没有获取到有效值时为空。
	IssueTime *string `json:"issue_time,omitempty"`

	// 证书生效时间，没有获取到有效值时为空。
	NotBefore *string `json:"not_before,omitempty"`

	// 证书失效时间，没有获取到有效值时为空。
	NotAfter *string `json:"not_after,omitempty"`

	// 证书有效期，按月为单位。[云证书管理服务提供了一种购买多年有效期证书的解决方案，这种多年有效期证书生效方式为多张有效期为1年的SSL证书叠加生效，例如：有效期为“3年”的证书实际包含3张有效期为1年且规格相同的SSL证书，在第一张证书到期前30天，系统自动以第一张证书的信息申请第二张证书，在第二张证书到期前30天，系统自动以第一张证书的信息申请第三张证书。](tag:hws)
	ValidityPeriod *int32 `json:"validity_period,omitempty"`

	// 域名认证方式，取值如下：DNS、FILE、EMAIL。
	ValidationMethod *string `json:"validation_method,omitempty"`

	// 域名类型，取值如下： - SINGLE_DOMAIN：单域名 - WILDCARD：通配符 - MULTI_DOMAIN：多域名
	DomainType *string `json:"domain_type,omitempty"`

	// 多域名类型，取值如下： - primary_single 主单 - primary_wildcard 主泛
	MultiDomainType *string `json:"multi_domain_type,omitempty"`

	// 证书绑定域名。
	Domain *string `json:"domain,omitempty"`

	// 证书绑定的附加域名信息。
	Sans *string `json:"sans,omitempty"`

	// 证书可绑定域名个数。
	DomainCount *int32 `json:"domain_count,omitempty"`

	// 证书可绑定附加域名个数。
	WildcardCount *int32 `json:"wildcard_count,omitempty"`

	// 证书的SHA-1指纹。
	Fingerprint *string `json:"fingerprint,omitempty"`

	// 企业项目ID，默认为“0”。 对于开通企业项目的用户，表示资源处于默认企业项目下。 对于未开通企业项目的用户，表示资源未处于企业项目下。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名所有权认证信息，详情请参见Authentification字段数据结构说明。
	Authentification *[]Authentification `json:"authentification,omitempty"`
	HttpStatusCode   int                 `json:"-"`
}

func (o ShowCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateResponse", string(data)}, " ")
}
