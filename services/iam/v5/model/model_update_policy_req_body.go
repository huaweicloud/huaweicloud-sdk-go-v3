package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePolicyReqBody **参数解释**： 更新自定义身份策略请求体。  **约束限制**： 不涉及。  **取值范围**： 不涉及。  **默认取值**： 不涉及。
type UpdatePolicyReqBody struct {

	// **参数解释**： 身份策略描述信息。  **约束限制**： 不涉及。  **取值范围**： 长度范围为[0,1000]。  **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o UpdatePolicyReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyReqBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyReqBody", string(data)}, " ")
}
