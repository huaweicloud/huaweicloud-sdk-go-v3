package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryVo struct {

	// 查询过滤器
	Filter *[]map[string]ConditionVo `json:"filter,omitempty"`

	// 过滤模式
	FilterMode *string `json:"filter_mode,omitempty"`

	Page *PageInfoVo `json:"page,omitempty"`

	// 排序条件
	Sort *[]SortInfo `json:"sort,omitempty"`

	// 返回字段
	ReturnFields *[]string `json:"return_fields,omitempty"`
}

func (o QueryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVo struct{}"
	}

	return strings.Join([]string{"QueryVo", string(data)}, " ")
}
