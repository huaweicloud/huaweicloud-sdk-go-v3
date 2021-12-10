package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AomMappingRuleResp struct {
	// 项目id

	ProjectId string `json:"project_id"`
	// 接入规则名称

	RuleName string `json:"rule_name"`
	// 接入规则id

	RuleId string `json:"rule_id"`

	RuleInfo *AomMappingRuleInfo `json:"rule_info"`
}

func (o AomMappingRuleResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AomMappingRuleResp struct{}"
	}

	return strings.Join([]string{"AomMappingRuleResp", string(data)}, " ")
}
