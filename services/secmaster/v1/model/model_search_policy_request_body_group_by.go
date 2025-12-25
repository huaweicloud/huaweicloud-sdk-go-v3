package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPolicyRequestBodyGroupBy 聚合条件
type SearchPolicyRequestBodyGroupBy struct {

	// 聚合字段
	GroupByFields *[]string `json:"group_by_fields,omitempty"`

	GroupByHit *SearchPolicyRequestBodyGroupByGroupByHit `json:"group_by_hit,omitempty"`
}

func (o SearchPolicyRequestBodyGroupBy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRequestBodyGroupBy struct{}"
	}

	return strings.Join([]string{"SearchPolicyRequestBodyGroupBy", string(data)}, " ")
}
