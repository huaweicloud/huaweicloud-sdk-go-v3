package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GetTemplatesRequest struct {
	// 模板平台类型

	Platform *string `json:"platform,omitempty"`
	// 语言类型

	Language *string `json:"language,omitempty"`
	// 是否支持流水线

	Pipeline *string `json:"pipeline,omitempty"`
	// 模板分类

	Entertype *string `json:"entertype,omitempty"`
	// 模板名称

	Search *string `json:"search,omitempty"`
	// 模板日期排序

	Dateorder *string `json:"dateorder,omitempty"`
	// 模板引用次数排序

	Usedtimeorder *string `json:"usedtimeorder,omitempty"`
	// 模板公开类型

	Type *string `json:"type,omitempty"`
	// 大区名称

	Region *string `json:"region,omitempty"`
	// 分页页数

	PageNo int32 `json:"page_no"`
	// 每页数据数

	PageSize int32 `json:"page_size"`
}

func (o GetTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTemplatesRequest struct{}"
	}

	return strings.Join([]string{"GetTemplatesRequest", string(data)}, " ")
}
