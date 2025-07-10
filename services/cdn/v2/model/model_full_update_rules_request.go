package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FullUpdateRulesRequest **参数解释：** 全量更新规则配置 **约束限制：** 不涉及
type FullUpdateRulesRequest struct {

	// **参数解释：** 规则配置内容，可以配置多个规则 **约束限制：** 不涉及
	Rules []CreateRuleRequest `json:"rules"`
}

func (o FullUpdateRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullUpdateRulesRequest struct{}"
	}

	return strings.Join([]string{"FullUpdateRulesRequest", string(data)}, " ")
}
