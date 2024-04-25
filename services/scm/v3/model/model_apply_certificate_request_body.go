package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplyCertificateRequestBody struct {

	// 该证书绑定的域名。 - 当购买的证书为“单域名”或“泛域名”类型的证书时，请直接填写单域名或泛域名即可。 - 当购买的证书为“多域名”类型的证书时，需要选择1个域名作为主域名。 示例：www.example.com
	Domain string `json:"domain"`

	// 绑定多域名类型证书的附加域名。 当购买的证书为“多域名”类型的证书，且有可增加附加域名的额度时，才需要设置该值。 多个域名需要以“;”隔开。 示例：www.example.com;www.example1.com;www.example2.com
	Sans *string `json:"sans,omitempty"`

	// 证书CSR串，与域名必须匹配。
	Csr *string `json:"csr,omitempty"`

	// 公司名称，OV和EV型证书必填。字符长度为0~63位。
	CompanyName *string `json:"company_name,omitempty"`

	// 部门名称。字符长度为0~63位。
	CompanyUnit *string `json:"company_unit,omitempty"`

	// 公司所在省份，OV和EV型证书必填。字符长度为0~63位。
	CompanyProvince *string `json:"company_province,omitempty"`

	// 公司所在市区，OV和EV型证书必填。字符长度为0~63位。
	CompanyCity *string `json:"company_city,omitempty"`

	// OV和EV型证书必填,国家编码，需符合正则\"**[A-Za-z]{2}**\"。
	Country *string `json:"country,omitempty"`

	// 申请人的姓名。请输入中文、英文字符，下划线，中划线，英文逗号，英文句点，且长度为4到100字节。
	ApplicantName string `json:"applicant_name"`

	// 申请人的电话号码。示例：13212345678
	ApplicantPhone string `json:"applicant_phone"`

	// 申请人的邮箱。示例：example@huawei.com
	ApplicantEmail string `json:"applicant_email"`

	// 技术联系人的姓名。字符长度为0~63位。
	ContactName *string `json:"contact_name,omitempty"`

	// 技术联系人的电话号码。示例：13212345678
	ContactPhone *string `json:"contact_phone,omitempty"`

	// 技术联系人的邮箱。示例：example@huawei.com
	ContactEmail *string `json:"contact_email,omitempty"`

	// 是否将DNS验证信息推送到华为云解析服务。 - true：推送。 - false：不推送。
	AutoDnsAuth *bool `json:"auto_dns_auth,omitempty"`

	// 是否同意授权隐私协议。此处仅能设置为true才能成功申请证书。 - true：同意隐私协议。 - false：不同意隐私协议。
	AgreePrivacyProtection bool `json:"agree_privacy_protection"`

	// 域名验证方式。 - DNS: DNS验证，指在域名管理平台通过解析指定的DNS记录，验证域名所有权。 - FILE: 文件验证，指通过在服务器上创建指定文件的方式来验证域名所有权。 - EMAIL: 邮箱验证，指登录域名管理员邮箱，接收域名确认邮件并根据提示进行操作来验证域名所有权。 DV域名型和DV基础版证书（GeoTrust入门级SSL证书和DigiCert免费SSL证书）默认通过“DNS验证”方式进行验证。 纯IP（公网IP）的证书仅支持通过“文件验证”方式进行验证，且仅纯IP证书支持“文件验证”方式验证。
	DomainMethod string `json:"domain_method"`

	// 密钥算法。默认RSA_2048
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`

	// 签名算法。Geo OV证书必填 - DEFAULT - SHA-256
	CaHashAlgorithm *string `json:"ca_hash_algorithm,omitempty"`
}

func (o ApplyCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ApplyCertificateRequestBody", string(data)}, " ")
}
