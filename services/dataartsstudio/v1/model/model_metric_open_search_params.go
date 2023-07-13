package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricOpenSearchParams struct {

	// 指标资产ID
	ArchitectureGuid *string `json:"architecture_guid,omitempty"`

	// 查询条件
	Query string `json:"query"`

	// 单次请求条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 是否按名称和描述搜索
	SearchNameDescription *bool `json:"search_name_description,omitempty"`

	// 是否查询子指标
	IncludeSubArchitecture *bool `json:"include_sub_architecture,omitempty"`
}

func (o MetricOpenSearchParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricOpenSearchParams struct{}"
	}

	return strings.Join([]string{"MetricOpenSearchParams", string(data)}, " ")
}
