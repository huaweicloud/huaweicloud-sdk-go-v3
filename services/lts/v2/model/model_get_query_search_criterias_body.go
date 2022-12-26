package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetQuerySearchCriteriasBody struct {

	// 快速查询字段
	Criteria *string `json:"criteria,omitempty"`

	// 快速查询名称
	Name *string `json:"name,omitempty"`

	// 快速查询id
	Id *string `json:"id,omitempty"`

	// 快速查询类型： 原始日志：ORIGINALLOG 可视化日志: VISUALIZATION
	SearchType *string `json:"search_type,omitempty"`
}

func (o GetQuerySearchCriteriasBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetQuerySearchCriteriasBody struct{}"
	}

	return strings.Join([]string{"GetQuerySearchCriteriasBody", string(data)}, " ")
}
