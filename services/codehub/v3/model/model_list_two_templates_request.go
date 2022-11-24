package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTwoTemplatesRequest struct {

	// 模板平台类型
	Platform *string `json:"platform,omitempty"`

	// 语言类型
	Language *string `json:"language,omitempty"`

	// 是否支持流水线
	Pipeline *string `json:"pipeline,omitempty"`

	// 模板分类
	EnterType *string `json:"enter_type,omitempty"`

	// 模板名称
	Search *string `json:"search,omitempty"`

	// 模板日期排序
	DateOrder *string `json:"date_order,omitempty"`

	// 模板引用次数排序
	UsedTimeOrder *string `json:"used_time_order,omitempty"`

	// 模板公开类型
	Type *string `json:"type,omitempty"`

	// 大区名称
	Region *string `json:"region,omitempty"`

	// 分页页数
	PageNo int32 `json:"page_no"`

	// 每页数据数
	PageSize int32 `json:"page_size"`
}

func (o ListTwoTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTwoTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTwoTemplatesRequest", string(data)}, " ")
}
