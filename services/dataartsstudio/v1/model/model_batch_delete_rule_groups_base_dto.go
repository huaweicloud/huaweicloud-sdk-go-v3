package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteRuleGroupsBaseDto struct {

	// 识别规则组id列表
	RuleGroupIds []string `json:"rule_group_ids"`
}

func (o BatchDeleteRuleGroupsBaseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteRuleGroupsBaseDto struct{}"
	}

	return strings.Join([]string{"BatchDeleteRuleGroupsBaseDto", string(data)}, " ")
}
