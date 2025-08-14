package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdpCertificate struct {

	// 证书全局唯一标识符（ID）
	CertificateId string `json:"certificate_id"`

	// 身份提供商签发者
	IssuerName string `json:"issuer_name"`

	// 证书有效期
	NotAfter float32 `json:"not_after"`

	// 证书有效期
	NotBefore float32 `json:"not_before"`

	// 证书公钥
	PublicKey string `json:"public_key"`

	// 证书序列号
	SerialNumber float32 `json:"serial_number"`

	// 证书序列号文本
	SerialNumberString string `json:"serial_number_string"`

	// 签名算法
	SignatureAlgorithmName string `json:"signature_algorithm_name"`

	// Subject
	SubjectName string `json:"subject_name"`

	// 版本
	Version float32 `json:"version"`

	// x509格式证书
	X509CertificateInPem string `json:"x509_Certificate_in_pem"`
}

func (o IdpCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpCertificate struct{}"
	}

	return strings.Join([]string{"IdpCertificate", string(data)}, " ")
}
