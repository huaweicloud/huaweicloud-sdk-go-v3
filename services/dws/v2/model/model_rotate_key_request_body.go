package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RotateKeyRequestBody **参数解释**： 轮转密钥请求体。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type RotateKeyRequestBody struct {

	// **参数解释**： KMS密钥ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MasterKeyId *string `json:"master_key_id,omitempty"`
}

func (o RotateKeyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RotateKeyRequestBody struct{}"
	}

	return strings.Join([]string{"RotateKeyRequestBody", string(data)}, " ")
}
