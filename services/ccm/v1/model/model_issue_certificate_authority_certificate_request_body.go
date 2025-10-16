package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueCertificateAuthorityCertificateRequestBody struct {

	// 父CA证书ID。
	IssuerId string `json:"issuer_id"`

	// 路径长度。
	PathLength *int32 `json:"path_length,omitempty"`

	// 签名哈希算法，可选值如下：   - **SHA256**   - **SHA384**   - **SHA512**   - **SM3**（中国站）
	SignatureAlgorithm string `json:"signature_algorithm"`

	Validity *Validity `json:"validity"`

	// 创建的CA类型。（如果激活包周期CA，为必填值） - **ROOT** : 根CA - **SUBORDINATE** : 从属CA
	Type *string `json:"type,omitempty"`

	DistinguishedName *DistinguishedName `json:"distinguished_name,omitempty"`

	// 密钥算法（如果激活包周期CA，为必填值），可选值如下：   - **RSA2048** : RSA算法，密钥长度2048位；   - **RSA4096** : RSA算法，密钥长度4096位；   - **EC256** : 椭圆曲线算法（Elliptic Curve Digital Signature Algorithm (ECDSA)），密钥长度256位；   - **EC384** : 椭圆曲线算法（Elliptic Curve Digital Signature Algorithm (ECDSA)），密钥长度384位；   - **SM2** : 国家密码管理局颁发的椭圆曲线算法（签名哈希算法SM3），密钥长度256位。（中国站）
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`

	// 密钥用法，具体标准参见RFC 5280中:[4.2.1.3节](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3)。   - **digitalSignature** : 数字签名；   - **nonRepudiation** : 不可抵赖；   - **keyEncipherment** : 密钥用于加密密钥数据；   - **dataEncipherment** : 用于加密数据；   - **keyAgreement** : 密钥协商；   - **keyCertSign** : 签发证书；   - **cRLSign** : 签发吊销列表；   - **encipherOnly** : 仅用于加密；   - **decipherOnly** : 仅用于解密。 > 缺省值如下： > - 根CA证书：默认为**[digitalSignature,keyCertSign,cRLSign]**，忽略用户传入值； > - 从属CA证书：默认为**[digitalSignature,keyCertSign,cRLSign]**，支持用户自定义。
	KeyUsages *[]string `json:"key_usages,omitempty"`

	CrlConfiguration *CrlConfiguration `json:"crl_configuration,omitempty"`

	// 企业多项目ID。用户未开通企业多项目时，不需要输入该字段。 用户开通企业多项目时，查询资源可以输入该字段。 若用户不输入该字段，默认查询租户所有有权限的企业多项目下的资源。 此时“enterprise_project_id”取值为“all”。 若用户输入该字段，取值满足以下任一条件.   取值为“all”   取值为“0”   满足正则匹配：“^[0-9a-z]{8}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{4}-[0-9a-z]{12}$”
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o IssueCertificateAuthorityCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCertificateAuthorityCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"IssueCertificateAuthorityCertificateRequestBody", string(data)}, " ")
}
