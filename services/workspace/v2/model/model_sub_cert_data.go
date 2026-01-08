package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubCertData 证书元数据。
type SubCertData struct {
	DistinguishedName *DistinguishedName `json:"distinguished_name"`

	// 密钥对生成算法 RSA-2048 RSA-3072。
	KeyAlgorithm string `json:"key_algorithm"`

	// 签名哈希算法 SHA-256 SHA-512。
	SignatureAlgorithm string `json:"signature_algorithm"`
}

func (o SubCertData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubCertData struct{}"
	}

	return strings.Join([]string{"SubCertData", string(data)}, " ")
}
