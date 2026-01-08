package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateReq 创建证书请求体。
type CreateCertificateReq struct {
	DistinguishedName *DistinguishedName `json:"distinguished_name"`

	// 密钥对生成算法 RSA-2048 RSA-3072。
	KeyAlgorithm string `json:"key_algorithm"`

	// 签名哈希算法 SHA-256 SHA-512。
	SignatureAlgorithm string `json:"signature_algorithm"`

	Validity *CertValidityData `json:"validity"`

	CrlConfiguration *CrlConfigurationData `json:"crl_configuration"`
}

func (o CreateCertificateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateReq struct{}"
	}

	return strings.Join([]string{"CreateCertificateReq", string(data)}, " ")
}
