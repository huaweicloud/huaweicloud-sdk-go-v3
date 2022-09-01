package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRulesRequest struct {

	// 规则对应的语言
	RuleLanguages *string `json:"rule_languages,omitempty" xml:"rule_languages"`

	// 缺陷等级，0致命，1严重，2一般，3提示
	RuleSeverity *string `json:"rule_severity,omitempty" xml:"rule_severity"`

	// 分页索引，偏移量
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRulesRequest struct{}"
	}

	return strings.Join([]string{"ListRulesRequest", string(data)}, " ")
}
