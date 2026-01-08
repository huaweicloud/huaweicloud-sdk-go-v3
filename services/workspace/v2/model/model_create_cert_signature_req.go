package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertSignatureReq 创建证书csr请求体。
type CreateCertSignatureReq struct {
	DistinguishedName *DistinguishedName `json:"distinguished_name"`

	// 密钥对生成算法 RSA-2048 RSA-3072。
	KeyAlgorithm string `json:"key_algorithm"`

	// 签名哈希算法 SHA-256 SHA-512。
	SignatureAlgorithm string `json:"signature_algorithm"`

	// 事务id。
	TransactionId *string `json:"transaction_id,omitempty"`
}

func (o CreateCertSignatureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertSignatureReq struct{}"
	}

	return strings.Join([]string{"CreateCertSignatureReq", string(data)}, " ")
}
