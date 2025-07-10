package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateRulesRequest **参数解释：** 批量更新规则状态和优先级 **约束限制：** 不涉及
type BatchUpdateRulesRequest struct {

	// **参数解释：** 规则列表 **约束限制：** 不涉及
	Rules []UpdateRuleStatusRequest `json:"rules"`
}

func (o BatchUpdateRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateRulesRequest", string(data)}, " ")
}
