package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAomMappingRulesResponse struct {

	// 项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// 接入规则名称
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name"`

	// 接入规则id
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	RuleInfo       *AomMappingRuleInfo `json:"rule_info,omitempty" xml:"rule_info"`
	HttpStatusCode int                 `json:"-"`
}

func (o UpdateAomMappingRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAomMappingRulesResponse struct{}"
	}

	return strings.Join([]string{"UpdateAomMappingRulesResponse", string(data)}, " ")
}
