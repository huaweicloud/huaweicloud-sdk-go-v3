package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchModifyReadAndWriteStrategyRequest struct {

	// **参数解释**：  读写策略。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReadWeightList *[]map[string]int32 `json:"read_weight_list,omitempty"`
}

func (o BatchModifyReadAndWriteStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyReadAndWriteStrategyRequest struct{}"
	}

	return strings.Join([]string{"BatchModifyReadAndWriteStrategyRequest", string(data)}, " ")
}
