package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTemplatesRequest Request Object
type ListTemplatesRequest struct {

	// 模板类型
	TemplateType string `json:"template_type"`

	// 是否内置模板
	IsBuildIn string `json:"is_build_in"`

	// 偏移量,表示从此偏移量开始查询,offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 模板名称，匹配规则为模糊匹配
	Name *string `json:"name,omitempty"`

	// 排序字段
	Sort *string `json:"sort,omitempty"`

	// 是否正序
	Asc *string `json:"asc,omitempty"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
