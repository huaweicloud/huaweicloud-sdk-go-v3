package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTwoTemplatesRequest struct {

	// 模板平台类型
	Platform *string `json:"platform,omitempty" xml:"platform"`

	// 语言类型
	Language *string `json:"language,omitempty" xml:"language"`

	// 是否支持流水线
	Pipeline *string `json:"pipeline,omitempty" xml:"pipeline"`

	// 模板分类
	EnterType *string `json:"enter_type,omitempty" xml:"enter_type"`

	// 模板名称
	Search *string `json:"search,omitempty" xml:"search"`

	// 模板日期排序
	DateOrder *string `json:"date_order,omitempty" xml:"date_order"`

	// 模板引用次数排序
	UsedTimeOrder *string `json:"used_time_order,omitempty" xml:"used_time_order"`

	// 模板公开类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 大区名称
	Region *string `json:"region,omitempty" xml:"region"`

	// 分页页数
	PageNo int32 `json:"page_no" xml:"page_no"`

	// 每页数据数
	PageSize int32 `json:"page_size" xml:"page_size"`
}

func (o ListTwoTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTwoTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTwoTemplatesRequest", string(data)}, " ")
}
