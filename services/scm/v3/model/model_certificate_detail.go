package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertificateDetail struct {

	// 证书id。
	Id string `json:"id"`

	// 证书名称。
	Name string `json:"name"`

	// 证书绑定的域名。
	Domain string `json:"domain"`

	// 多域名证书绑定的附加域名。
	Sans string `json:"sans"`

	// 签名算法。
	SignatureAlgorithm string `json:"signature_algorithm"`

	// 是否支持部署。
	DeploySupport bool `json:"deploy_support"`

	// 证书类型。取值如下： DV_SSL_CERT、DV_SSL_CERT_BASIC、EV_SSL_CERT、 EV_SSL_CERT_PRO、OV_SSL_CERT、OV_SSL_CERT_PRO
	Type string `json:"type"`

	// 证书品牌。取值如下：GLOBALSIGN、SYMANTEC、GEOTRUST、CFCA
	Brand string `json:"brand"`

	// 证书过期时间。
	ExpireTime string `json:"expire_time"`

	// 域名类型。取值如下： - SINGLE_DOMAIN：单域名 - WILDCARD：通配符 - MULTI_DOMAIN：多域名
	DomainType string `json:"domain_type"`

	// 证书有效期，按月为单位。[云证书管理服务提供了一种购买多年有效期证书的解决方案，这种多年有效期证书生效方式为多张有效期为1年的SSL证书叠加生效，例如：有效期为“3年”的证书实际包含3张有效期为1年且规格相同的SSL证书，在第一张证书到期前30天，系统自动以第一张证书的信息申请第二张证书，在第二张证书到期前30天，系统自动以第一张证书的信息申请第三张证书。](tag:hws)
	ValidityPeriod int32 `json:"validity_period"`

	// 证书状态，取值如下： - PAID：证书已支付；待申请证书。 - ISSUED：证书已签发。 - CHECKING：证书申请审核中。 - CANCELCHECKING：取消证书申请审核中。 - UNPASSED：证书申请未通过。 - EXPIRED：证书已过期。 - REVOKING：证书吊销申请审核中。 - CANCLEREVOKING：证书取消吊销申请审核中。 - REVOKED：证书已吊销。 - UPLOAD：证书托管中。 - SUPPLEMENTCHECKING：多域名证书新增附加域名审核中。 - CANCELSUPPLEMENTING：取消新增附加域名审核中。
	Status string `json:"status"`

	// 证书可绑定域名个数。
	DomainCount int32 `json:"domain_count"`

	// 证书可绑定泛域名个数。
	WildcardCount int32 `json:"wildcard_count"`

	// 证书描述。
	Description string `json:"description"`

	// 企业项目ID，默认为“0”。 对于开通企业项目的用户，表示资源处于默认企业项目下。 对于未开通企业项目的用户，表示资源未处于企业项目下。
	EnterpriseProjectId string `json:"enterprise_project_id"`
}

func (o CertificateDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertificateDetail struct{}"
	}

	return strings.Join([]string{"CertificateDetail", string(data)}, " ")
}
