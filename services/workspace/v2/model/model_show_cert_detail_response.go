package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertDetailResponse Response Object
type ShowCertDetailResponse struct {

	// 证书id。
	CertId *string `json:"cert_id,omitempty"`

	// 序列号。
	SerialNumber *string `json:"serial_number,omitempty"`

	// 证书类型。
	Type *string `json:"type,omitempty"`

	// 证书应用范围。
	Apply *string `json:"apply,omitempty"`

	DistinguishedName *DistinguishedName `json:"distinguished_name,omitempty"`

	// 密钥对生成算法 RSA-2048 RSA-3072。
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`

	// 签名哈希算法 SHA-256 SHA-512。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 生效时间。
	NotBefore *string `json:"not_before,omitempty"`

	// 过期时间。
	NotAfter *string `json:"not_after,omitempty"`

	// 证书状态 DISABLE,ENABLE,EXPIRED,DELETE。
	Status *string `json:"status,omitempty"`

	// 证书pem编码。
	PemCode *string `json:"pem_code,omitempty"`

	// 颁发者名字。
	IssuerName *string `json:"issuer_name,omitempty"`

	CrlConfiguration *CrlConfigurationData `json:"crl_configuration,omitempty"`
	HttpStatusCode   int                   `json:"-"`
}

func (o ShowCertDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowCertDetailResponse", string(data)}, " ")
}
