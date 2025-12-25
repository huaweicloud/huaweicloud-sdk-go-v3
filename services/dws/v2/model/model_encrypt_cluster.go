package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncryptCluster **参数解释**： 转加密集群请求信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type EncryptCluster struct {

	// **参数解释**： KMS密钥ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MasterKeyId string `json:"master_key_id"`

	// **参数解释**： KMS密钥名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MasterKeyName *string `json:"master_key_name,omitempty"`

	// **参数解释**： 加密类型。国密、国际加密。 **约束限制**： 不涉及。 **取值范围**： - generalCipher：AES-CBC算法加密。 - SMcompatible：sm4算法加密。  **默认取值**： 不涉及。
	CryptAlgorithm string `json:"crypt_algorithm"`
}

func (o EncryptCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncryptCluster struct{}"
	}

	return strings.Join([]string{"EncryptCluster", string(data)}, " ")
}
