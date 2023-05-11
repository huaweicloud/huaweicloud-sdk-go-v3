package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateAomMappingRequest struct {

	// 接入规则id
	RuleId string `json:"rule_id"`

	// 项目id
	ProjectId string `json:"project_id"`

	// 接入规则名称
	RuleName string `json:"rule_name"`

	RuleInfo *AomMappingRuleInfo `json:"rule_info"`
}

func (o UpdateAomMappingRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAomMappingRequest struct{}"
	}

	return strings.Join([]string{"UpdateAomMappingRequest", string(data)}, " ")
}
