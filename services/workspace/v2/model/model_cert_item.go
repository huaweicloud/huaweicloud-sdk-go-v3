package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CertItem 证书信息。
type CertItem struct {

	// 租户id。
	TenantId *string `json:"tenant_id,omitempty"`

	// 证书id。
	CertId *string `json:"cert_id,omitempty"`

	// 证书名。
	CommonName *string `json:"common_name,omitempty"`

	// 证书类型ROOT, SUBORDINATE。
	Type *string `json:"type,omitempty"`

	// 证书状态 DISABLE,ENABLE,EXPIRED,DELETE。
	Status *string `json:"status,omitempty"`

	// 密钥生成算法 RSA-2048,RSA-3072。
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`

	// 签名哈希算法。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 应用场景。
	Apply *string `json:"apply,omitempty"`

	// 生效时间。
	NotBefore *string `json:"not_before,omitempty"`

	// 过期时间。
	NotAfter *string `json:"not_after,omitempty"`
}

func (o CertItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertItem struct{}"
	}

	return strings.Join([]string{"CertItem", string(data)}, " ")
}
