package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DecryptDatakeyCapsuleRequestBody struct {

	// **参数解释：** 密钥ID **约束限制：** UUID格式，满足正则表达式^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$ **取值范围：** 不涉及 **默认取值：** 不涉及
	KeyId string `json:"key_id"`

	// **参数解释：** 公钥信息，使用RSAES_OAEP_SHA_256算法加密；如果传递了public_key，KMS会使用该公钥对明文数据密钥进行加密，并返回加密后的数据密钥 **约束限制：** 仅支持RSA公钥 **取值范围：** 不涉及 **默认取值：** 不涉及
	PublicKey *string `json:"public_key,omitempty"`

	// **参数解释：** 密钥胶囊 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	DatakeyCapsule string `json:"datakey_capsule"`

	AttestationDocument *DecryptDatakeyCapsuleRequestBodyAttestationDocument `json:"attestation_document"`
}

func (o DecryptDatakeyCapsuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyCapsuleRequestBody struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyCapsuleRequestBody", string(data)}, " ")
}
