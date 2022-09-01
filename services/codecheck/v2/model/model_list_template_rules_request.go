package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplateRulesRequest struct {

	// 项目ID
	ProjectId string `json:"project_id" xml:"project_id"`

	// 规则集ID
	RulesetId string `json:"ruleset_id" xml:"ruleset_id"`

	// 规则状态  '1查询全部，2已启动，3未启用'
	Types string `json:"types" xml:"types"`

	// 规则语言
	Languages *string `json:"languages,omitempty" xml:"languages"`

	// 规则所属标签
	Tags *string `json:"tags,omitempty" xml:"tags"`

	// 分页索引，偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListTemplateRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateRulesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplateRulesRequest", string(data)}, " ")
}
