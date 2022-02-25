package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SetDefaulTemplateRequest struct {
	// 项目ID

	ProjectId string `json:"project_id"`
	// 规则集ID

	RulesetId string `json:"ruleset_id"`
	// 规则集语言

	Language string `json:"language"`
}

func (o SetDefaulTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDefaulTemplateRequest struct{}"
	}

	return strings.Join([]string{"SetDefaulTemplateRequest", string(data)}, " ")
}
