package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TemplateQueryV2 struct {

	// 模板分类数组。
	Category *[]string `json:"category,omitempty"`

	// 搜索关键字，支持按名称和描述搜索，默认null。
	Keyword *string `json:"keyword,omitempty"`

	// 排序字段和排序顺序指定。比如： - desc(created_at)：根据创建时间降序 - desc(usage_count)：根据引用次数降序
	SortBy *string `json:"sort_by,omitempty"`

	// 标签： - all：全部 - new：最新 - hot：热门 - recommend：推荐
	Label *string `json:"label,omitempty"`

	// 是否查询用户自己创建的模板，默认查所有模板。
	MyTemplates *bool `json:"my_templates,omitempty"`

	// 查所有模板时只处理上架的；查用户模板，需支持按状态查询，状态： - 0：审核中 - 1：上架 - 2：下架 不传表示查所有的（默认）
	Status *int32 `json:"status,omitempty"`

	// 模板状态数组。
	StatusArray *[]int32 `json:"status_array,omitempty"`

	// 是否查询有消息的模板，默认查所有模板。
	HasNotices *bool `json:"has_notices,omitempty"`

	// 模板关联的云产品(产品短名)列表。
	Productshorts *[]string `json:"productshorts,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页的模板条数。
	Limit *int32 `json:"limit,omitempty"`

	// 模板关联的自定义标签列表。
	TagIds *[]string `json:"tag_ids,omitempty"`

	// 模板类型： - 0：doc - 1：code - 2：pipeline - 3：devops
	Types *[]int32 `json:"types,omitempty"`

	// 动、静态代码模板标识： - 0：动态模板codetemplate - 1：静态模板codesample
	IsStatic *int32 `json:"is_static,omitempty"`

	// 平台来源： - 0：codelabs - 1：devstar
	PlatformSource *[]int32 `json:"platform_source,omitempty"`

	// 模板关联的标签名称列表。
	TagNames *[]string `json:"tag_names,omitempty"`
}

func (o TemplateQueryV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateQueryV2 struct{}"
	}

	return strings.Join([]string{"TemplateQueryV2", string(data)}, " ")
}
