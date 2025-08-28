package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowL7PolicyRequest Request Object
type ShowL7PolicyRequest struct {

	// **参数解释**：转发策略ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	L7policyId string `json:"l7policy_id"`
}

func (o ShowL7PolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowL7PolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowL7PolicyRequest", string(data)}, " ")
}
