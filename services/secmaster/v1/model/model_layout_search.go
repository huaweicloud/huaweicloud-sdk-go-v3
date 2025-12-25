package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayoutSearch 布局搜索条件
type LayoutSearch struct {

	// 布局名称
	Name *string `json:"name,omitempty"`

	// 布局类型
	UsedBy *string `json:"used_by,omitempty"`

	// 创建布局选择的数据类的business_code
	BindingCode *string `json:"binding_code,omitempty"`

	// 是否为系统内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 是否为模板
	IsTemplate *bool `json:"is_template,omitempty"`

	// 是否为默认布局
	IsDefault *bool `json:"is_default,omitempty"`

	// 布局页面类型
	LayoutType *string `json:"layout_type,omitempty"`

	// 分页
	Offset *int32 `json:"offset,omitempty"`

	// 每页个数
	Limit *int32 `json:"limit,omitempty"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 升序或者降序排列
	SortDir *string `json:"sort_dir,omitempty"`

	// 搜索关键字
	SearchTxt *string `json:"search_txt,omitempty"`

	// 起始时间
	FromDate *string `json:"from_date,omitempty"`

	// 结束时间
	ToDate *string `json:"to_date,omitempty"`
}

func (o LayoutSearch) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayoutSearch struct{}"
	}

	return strings.Join([]string{"LayoutSearch", string(data)}, " ")
}
