package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplateRulesRequest struct {
	// 规则集ID

	RulesetId string `json:"ruleset_id"`
	// 规则状态  '1查询全部，2已启动，3未启用'

	Types string `json:"types"`
	// 规则语言

	Languages *string `json:"languages,omitempty"`
	// 分页索引，偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTemplateRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateRulesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateRulesRequest", string(data)}, " ")
}
