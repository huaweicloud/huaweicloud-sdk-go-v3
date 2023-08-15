package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCertificateRequestBody struct {

	// 父CA证书ID。
	IssuerId string `json:"issuer_id"`

	// 密钥算法，可选值如下：   - **RSA2048** : RSA算法，密钥长度2048位；   - **RSA4096** : RSA算法，密钥长度4096位；   - **EC256** : 椭圆曲线算法（Elliptic Curve Digital Signature Algorithm (ECDSA)），密钥长度256位；   - **EC384** : 椭圆曲线算法（Elliptic Curve Digital Signature Algorithm (ECDSA)），密钥长度384位；   - **SM2** : 国家密码管理局颁发的椭圆曲线算法（签名哈希算法SM3），密钥长度256位。（中国站）
	KeyAlgorithm string `json:"key_algorithm"`

	// 签名哈希算法，可选值如下：   - **SHA256**   - **SHA384**   - **SHA512**   - **SM3**（中国站）
	SignatureAlgorithm string `json:"signature_algorithm"`

	DistinguishedName *CertDistinguishedName `json:"distinguished_name"`

	Validity *Validity `json:"validity"`

	// 密钥用法，具体标准参见RFC 5280中:[4.2.1.3节](https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.3)。   - **digitalSignature** : 数字签名；   - **nonRepudiation** : 不可抵赖；   - **keyEncipherment** : 密钥用于加密密钥数据；   - **dataEncipherment** : 用于加密数据；   - **keyAgreement** : 密钥协商；   - **keyCertSign** : 签发证书；   - **cRLSign** : 签发吊销列表；   - **encipherOnly** : 仅用于加密；   - **decipherOnly** : 仅用于解密。
	KeyUsages *[]string `json:"key_usages,omitempty"`

	// 主体备用名称，详情请参见**SubjectAlternativeName**字段数据结构说明。   - array大小：[0,20]。
	SubjectAlternativeNames *[]SubjectAlternativeName `json:"subject_alternative_names,omitempty"`

	ExtendedKeyUsage *ExtendedKeyUsage `json:"extended_key_usage,omitempty"`

	CustomizedExtension *CustomizedExtension `json:"customized_extension,omitempty"`
}

func (o CreateCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequestBody", string(data)}, " ")
}
