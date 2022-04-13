package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTemplateRulesResponse struct {
	// 规则集的规则列表信息

	Info *[]RuleItem `json:"info,omitempty"`
	// 总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTemplateRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplateRulesResponse struct{}"
	}

	return strings.Join([]string{"ListTemplateRulesResponse", string(data)}, " ")
}
