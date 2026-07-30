package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssociatedResourceRulesResponse Response Object
type ListAssociatedResourceRulesResponse struct {

	// 规则信息
	Rules *[]AssociatedResourceRule `json:"rules,omitempty"`

	// 记录总数
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListAssociatedResourceRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssociatedResourceRulesResponse struct{}"
	}

	return strings.Join([]string{"ListAssociatedResourceRulesResponse", string(data)}, " ")
}
