package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteRulesetRequest struct {

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 规则集ID
	RulesetId string `json:"ruleset_id" xml:"ruleset_id"`
}

func (o DeleteRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRulesetRequest struct{}"
	}

	return strings.Join([]string{"DeleteRulesetRequest", string(data)}, " ")
}
