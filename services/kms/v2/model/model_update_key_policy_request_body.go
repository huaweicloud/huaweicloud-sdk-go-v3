package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateKeyPolicyRequestBody struct {

	// **参数解释：** 修改的密钥策略 **约束限制：** 转义后的JSON字符串 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policy string `json:"policy"`

	// **参数解释：** 密钥策略描述信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`
}

func (o UpdateKeyPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateKeyPolicyRequestBody", string(data)}, " ")
}
