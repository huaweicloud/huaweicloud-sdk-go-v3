package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KmsInfo struct {

	// AES_256算法密钥数量
	Aes256 *int32 `json:"aes_256,omitempty"`

	// SM4算法密钥数量
	Sm4 *int32 `json:"sm4,omitempty"`

	// RSA_2048算法密钥数量
	Rsa2048 *int32 `json:"rsa_2048,omitempty"`

	// RSA_3072算法密钥数量
	Rsa3072 *int32 `json:"rsa_3072,omitempty"`

	// RSA_4096算法密钥数量
	Rsa4096 *int32 `json:"rsa_4096,omitempty"`

	// EC_P256算法密钥数量
	EcP256 *int32 `json:"ec_p256,omitempty"`

	// EC_P384务算法密钥数量
	EcP384 *int32 `json:"ec_p384,omitempty"`

	// SM2算法密钥数量
	Sm2 *int32 `json:"sm2,omitempty"`
}

func (o KmsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KmsInfo struct{}"
	}

	return strings.Join([]string{"KmsInfo", string(data)}, " ")
}
