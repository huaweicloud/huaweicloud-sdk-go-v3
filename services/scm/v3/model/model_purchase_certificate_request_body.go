package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PurchaseCertificateRequestBody struct {

	// 证书品牌，取值如下： - GEOTRUST - GLOBALSIGN - SYMANTEC - CFCA - TRUSTASIA - VTRUS
	CertBrand string `json:"cert_brand"`

	// 证书类型，取值如下： - DV_SSL_CERT - DV_SSL_CERT_BASIC - EV_SSL_CERT - EV_SSL_CERT_PRO - OV_SSL_CERT - OV_SSL_CERT_PRO
	CertType string `json:"cert_type"`

	// 域名类型，取值如下： - SINGLE_DOMAIN：单域名类型。 - MULTI_DOMAIN：多域名类型。 - WILDCARD：泛域名类型。
	DomainType string `json:"domain_type"`

	// 证书有效期（年）。
	EffectiveTime int32 `json:"effective_time"`

	// 域名数量。 - 当“domain_type”选择的是“SINGLE_DOMAIN”或“WILDCARD”类型的证书时，域名数量取值为“1”。 - 当“domain_type”选择的是“MULTI_DOMAIN”类型的证书时，域名数量取值范围为“2~100”。
	DomainNumbers int32 `json:"domain_numbers"`

	// 购买的证书数量。取值范围为1~100。
	OrderNumber int32 `json:"order_number"`

	// 是否同意隐私协议，此处仅能设置为true才能成功购买证书。 - true：同意隐私协议。 - false：不同意隐私协议。
	AgreePrivacyProtection bool `json:"agree_privacy_protection"`

	// 多域名中的主域名类型 - SINGLE_DOMAIN：主单域名 - WILDCARD_DOMAIN：主泛域名
	PrimaryDomainType *string `json:"primary_domain_type,omitempty"`

	// 附加单域名数量。
	SingleDomainNumber *int32 `json:"single_domain_number,omitempty"`

	// 附加泛域名数量。
	WildcardDomainNumber *int32 `json:"wildcard_domain_number,omitempty"`

	// 是否开启自动支付。 - true：开启自动支付。 - false：不开启自动支付。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// 企业多项目ID。用户未开通企业多项目时，不需要输入该字段。 用户开通企业多项目时，查询资源可以输入该字段。 若用户不输入该字段，默认查询租户所有有权限的企业多项目下的资源。 此时“enterprise_project_id”取值为“all”。 若用户输入该字段，取值满足以下任一条件.  取值为“all”  取值为“0”  满足正则匹配：“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 订单号。仅组合购场景使用
	OrderId *string `json:"order_id,omitempty"`
}

func (o PurchaseCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PurchaseCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"PurchaseCertificateRequestBody", string(data)}, " ")
}
