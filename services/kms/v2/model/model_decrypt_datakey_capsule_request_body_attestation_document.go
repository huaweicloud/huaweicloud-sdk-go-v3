package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DecryptDatakeyCapsuleRequestBodyAttestationDocument **参数解释：** 接入点证明文档 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type DecryptDatakeyCapsuleRequestBodyAttestationDocument struct {

	// **参数解释：** ECS证明文档 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EcsSignature *string `json:"ecs_signature,omitempty"`

	// **参数解释：** 通用类型接入点的签名信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CustomSignature *string `json:"custom_signature,omitempty"`

	// **参数解释：** 通用类型接入点公钥信息 **约束限制：** 格式是X509公钥格式中的Base64字符串 **取值范围：** 不涉及 **默认取值：** 不涉及
	CustomPublicKey *string `json:"custom_public_key,omitempty"`

	// **参数解释：** 通用类型签名信息过期时间 **约束限制：** 时间格式是ISO 8601格式，yyyy-mm-ddTHH:MM:SSZ **取值范围：** 不涉及 **默认取值：** 不涉及
	ExpireTime *string `json:"expire_time,omitempty"`

	// **参数解释：** CCE类型访问凭证 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ServiceToken *string `json:"service_token,omitempty"`
}

func (o DecryptDatakeyCapsuleRequestBodyAttestationDocument) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DecryptDatakeyCapsuleRequestBodyAttestationDocument struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyCapsuleRequestBodyAttestationDocument", string(data)}, " ")
}
