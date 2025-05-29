package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TagResourceResourceDetail 证书资源详情。
type TagResourceResourceDetail struct {

	// 证书ID。
	CertId *string `json:"cert_id,omitempty"`

	// 证书名称。字符长度为3~63位。
	CertName *string `json:"cert_name,omitempty"`

	// 该证书绑定的域名。 - 当购买的证书为“单域名”或“泛域名”类型的证书时，请直接填写单域名或泛域名即可。 - 当购买的证书为“多域名”类型的证书时，需要选择1个域名作为主域名。 示例：www.example.com
	Domain *string `json:"domain,omitempty"`

	// 证书类型。 - OV_SSL_CERT：企业型SSL证书。 - EV_SSL_CERT：增强型SSL证书。
	CertType *string `json:"cert_type,omitempty"`

	// 证书品牌。GLOBALSIGN：GlobalSign品牌。
	CertBrand *string `json:"cert_brand,omitempty"`

	// 域名类型。 - SINGLE_DOMAIN：单域名类型。 - MULTI_DOMAIN：多域名类型。 - WILDCARD：泛域名类型。
	DomainType *string `json:"domain_type,omitempty"`

	// 证书有效期（年）。
	PurchasePeriod *int32 `json:"purchase_period,omitempty"`

	// 证书到期时间，毫秒级时间戳。
	ExpiredTime *string `json:"expired_time,omitempty"`

	// 订单状态。
	OrderStatus *string `json:"order_status,omitempty"`

	// 域名数量。 - 当“domain_type”选择的是“SINGLE_DOMAIN”或“WILDCARD”类型的证书时，域名数量取值为“1”。 - 当“domain_type”选择的是“MULTI_DOMAIN”类型的证书时，域名数量取值范围为“2~250”。
	DomainNum *int32 `json:"domain_num,omitempty"`

	// 泛域名数量。
	WildcardNumber *int32 `json:"wildcard_number,omitempty"`

	// 证书绑定的附加域名信息。
	Sans *string `json:"sans,omitempty"`

	// 证书描述。
	CertDes *string `json:"cert_des,omitempty"`

	// 签名算法。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 失败原因。
	FailReason *string `json:"fail_reason,omitempty"`

	// 订单流水号。
	PartnerOrderId *string `json:"partner_order_id,omitempty"`

	// 证书是否支持推送。
	PushSupport *bool `json:"push_support,omitempty"`

	// 证书签发时间，毫秒级时间戳。
	CertIssuedTime *string `json:"cert_issued_time,omitempty"`

	// 资源id。
	ResourceId *string `json:"resource_id,omitempty"`

	// 证书是否支持退订。
	UnsubscribeSupport *bool `json:"unsubscribe_support,omitempty"`

	// 企业项目ID，默认为“0”。 对于开通企业项目的用户，表示资源处于默认企业项目下。 对于未开通企业项目的用户，表示资源未处于企业项目下。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 初始证书id。
	OriginCertId *string `json:"origin_cert_id,omitempty"`

	// 续费购买证书id。
	RenewalCertId *string `json:"renewal_cert_id,omitempty"`

	// 自动续费状态。
	AutoRenewStatus *int32 `json:"auto_renew_status,omitempty"`

	// 剩余证书有效个数。
	RemainCertNumber *int32 `json:"remain_cert_number,omitempty"`

	// 证书是否支持自动部署。
	AutoDeploySupport *bool `json:"auto_deploy_support,omitempty"`
}

func (o TagResourceResourceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceResourceDetail struct{}"
	}

	return strings.Join([]string{"TagResourceResourceDetail", string(data)}, " ")
}
