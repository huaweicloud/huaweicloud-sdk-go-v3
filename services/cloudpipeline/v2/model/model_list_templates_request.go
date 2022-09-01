package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListTemplatesRequest struct {

	// 语言类型 中文:zh-cn 英文:en-us，默认en-us
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 模板类型
	TemplateType string `json:"template_type" xml:"template_type"`

	// 是否内置模板
	IsBuildIn string `json:"is_build_in" xml:"is_build_in"`

	// 偏移量,表示从此偏移量开始查询,offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 模板名称，匹配规则为模糊匹配
	Name *string `json:"name,omitempty" xml:"name"`

	// 排序字段
	Sort *string `json:"sort,omitempty" xml:"sort"`

	// 是否正序
	Asc *string `json:"asc,omitempty" xml:"asc"`
}

func (o ListTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListTemplatesRequest", string(data)}, " ")
}
