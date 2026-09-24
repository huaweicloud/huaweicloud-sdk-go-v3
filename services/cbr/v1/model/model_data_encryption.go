package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataEncryption struct {

	// 存储库的密钥ID。如果为非加密存储库，默认值为None
	Cmkid *string `json:"cmkid,omitempty"`

	// 存储库的加密算法类型。如果为非加密存储库，默认值为None
	EncryptedAlgorithm *string `json:"encrypted_algorithm,omitempty"`
}

func (o DataEncryption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataEncryption struct{}"
	}

	return strings.Join([]string{"DataEncryption", string(data)}, " ")
}
