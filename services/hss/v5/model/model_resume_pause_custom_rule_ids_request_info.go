package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResumePauseCustomRuleIdsRequestInfo struct {

	// **参数解释**: 启用、停用 **约束限制**: 必填 **取值范围**: - 1：启用 - 0：停用  **默认取值**: 不涉及
	Enable int32 `json:"enable"`

	// **参数解释**： 规则ID列表 **约束限制**: 必填 **取值范围**: 1-1000个规则值 **默认取值**: 不涉及
	RuleIdList []string `json:"rule_id_list"`
}

func (o ResumePauseCustomRuleIdsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePauseCustomRuleIdsRequestInfo struct{}"
	}

	return strings.Join([]string{"ResumePauseCustomRuleIdsRequestInfo", string(data)}, " ")
}
