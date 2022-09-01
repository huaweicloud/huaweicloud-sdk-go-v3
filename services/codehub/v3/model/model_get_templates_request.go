package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type GetTemplatesRequest struct {

	// 模板平台类型
	Platform *string `json:"platform,omitempty" xml:"platform"`

	// 语言类型
	Language *string `json:"language,omitempty" xml:"language"`

	// 是否支持流水线
	Pipeline *string `json:"pipeline,omitempty" xml:"pipeline"`

	// 模板分类
	Entertype *string `json:"entertype,omitempty" xml:"entertype"`

	// 模板名称
	Search *string `json:"search,omitempty" xml:"search"`

	// 模板日期排序
	Dateorder *string `json:"dateorder,omitempty" xml:"dateorder"`

	// 模板引用次数排序
	Usedtimeorder *string `json:"usedtimeorder,omitempty" xml:"usedtimeorder"`

	// 模板公开类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 大区名称
	Region *string `json:"region,omitempty" xml:"region"`

	// 分页页数
	PageNo int32 `json:"page_no" xml:"page_no"`

	// 每页数据数
	PageSize int32 `json:"page_size" xml:"page_size"`
}

func (o GetTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetTemplatesRequest struct{}"
	}

	return strings.Join([]string{"GetTemplatesRequest", string(data)}, " ")
}
