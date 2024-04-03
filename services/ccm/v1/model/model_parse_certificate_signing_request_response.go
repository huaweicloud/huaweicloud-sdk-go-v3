package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseCertificateSigningRequestResponse Response Object
type ParseCertificateSigningRequestResponse struct {

	// 密钥算法。
	KeyAlgorithm *string `json:"key_algorithm,omitempty"`

	// 密钥算法长度，单位为bit。
	KeyAlgorithmLength *int32 `json:"key_algorithm_length,omitempty"`

	// 签名算法，带具体的签名与哈希算法，如\"SHA256withRSA\"。
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`

	// 公钥内容。 > 其中，换行符已被“\\r\\n”替代；
	PublicKey *string `json:"public_key,omitempty"`

	DistinguishedName *DistinguishedName `json:"distinguished_name,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ParseCertificateSigningRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseCertificateSigningRequestResponse struct{}"
	}

	return strings.Join([]string{"ParseCertificateSigningRequestResponse", string(data)}, " ")
}
