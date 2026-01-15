package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCertificatePemReq 导入pem证书。
type ImportCertificatePemReq struct {
	DistinguishedName *DistinguishedName `json:"distinguished_name"`

	// 密钥对生成算法 RSA-2048 RSA-3072。
	KeyAlgorithm string `json:"key_algorithm"`

	// 事务id。
	TransactionId *string `json:"transaction_id,omitempty"`

	CrlConfiguration *CrlConfigurationData `json:"crl_configuration,omitempty"`

	// 证书pem。
	PemCode *string `json:"pem_code,omitempty"`
}

func (o ImportCertificatePemReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificatePemReq struct{}"
	}

	return strings.Join([]string{"ImportCertificatePemReq", string(data)}, " ")
}
