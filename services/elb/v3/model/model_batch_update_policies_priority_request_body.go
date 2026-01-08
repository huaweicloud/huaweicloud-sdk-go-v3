package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdatePoliciesPriorityRequestBody This is a auto create Body Object
type BatchUpdatePoliciesPriorityRequestBody struct {

	// **参数解释**：批量更新转发策略优先级的请求参数。  **约束限制**：不涉及
	L7policies *[]BatchUpdatePriorityRequestBody `json:"l7policies,omitempty"`
}

func (o BatchUpdatePoliciesPriorityRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePoliciesPriorityRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdatePoliciesPriorityRequestBody", string(data)}, " ")
}
