package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteCustomRuleIdsRequestInfo struct {

	// **参数解释**： 规则ID列表 **约束限制**: 必填 **取值范围**: 1-1000个规则ID **默认取值**: 不涉及
	RuleIdList []string `json:"rule_id_list"`
}

func (o DeleteCustomRuleIdsRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCustomRuleIdsRequestInfo struct{}"
	}

	return strings.Join([]string{"DeleteCustomRuleIdsRequestInfo", string(data)}, " ")
}
