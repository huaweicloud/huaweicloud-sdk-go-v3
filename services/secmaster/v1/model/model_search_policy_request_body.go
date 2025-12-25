package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPolicyRequestBody 查询应急策略列表入参
type SearchPolicyRequestBody struct {

	// 分页数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	Condition *SearchPolicyRequestBodyCondition `json:"condition,omitempty"`

	// 排序条件
	Sort *[]SearchPolicyRequestBodySort `json:"sort,omitempty"`

	GroupBy *SearchPolicyRequestBodyGroupBy `json:"group_by,omitempty"`
}

func (o SearchPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SearchPolicyRequestBody", string(data)}, " ")
}
