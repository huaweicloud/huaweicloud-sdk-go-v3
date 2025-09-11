package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SearchIpdIssuesRequestBody struct {

	// 过滤条件
	Filter *[]map[string]ConditionVo `json:"filter,omitempty"`

	Page *PageInfoVo `json:"page,omitempty"`

	// 返回字段
	ReturnFields *[]string `json:"return_fields,omitempty"`

	// 排序字段
	Sort *[]SortInfo `json:"sort,omitempty"`
}

func (o SearchIpdIssuesRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchIpdIssuesRequestBody struct{}"
	}

	return strings.Join([]string{"SearchIpdIssuesRequestBody", string(data)}, " ")
}
