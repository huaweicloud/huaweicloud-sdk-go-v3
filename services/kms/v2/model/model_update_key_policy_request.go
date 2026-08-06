package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKeyPolicyRequest Request Object
type UpdateKeyPolicyRequest struct {

	// **参数解释：** 密钥策略ID。 **约束限制：** 不涉及 **取值范围：** UUID格式，字符长度36-36。 **默认取值：** 不涉及
	PolicyId string `json:"policy_id"`

	Body *UpdateKeyPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyPolicyRequest", string(data)}, " ")
}
