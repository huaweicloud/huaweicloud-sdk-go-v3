package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchParametersExt 请求参数
type SearchParametersExt struct {

	// 属性
	Attributes *interface{} `json:"attributes,omitempty"`

	// 分类
	Classifications *[]string `json:"classifications,omitempty"`

	// 数据连接
	ConnectionNames *[]string `json:"connection_names,omitempty"`

	// 是否排除分类
	ExcludeClassifications *bool `json:"exclude_classifications,omitempty"`

	// 是否排除密级
	ExcludeSecurityLevels *bool `json:"exclude_security_levels,omitempty"`

	// 是否排除标签
	ExcludeTags *bool `json:"exclude_tags,omitempty"`

	// 包含分类属性
	IncludeClassificationAttributes *bool `json:"include_classification_attributes,omitempty"`

	// 包含子分类
	IncludeSubClassifications *bool `json:"include_sub_classifications,omitempty"`

	// 分页参数，每页限制数量，默认查询所有
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数，偏移量，默认查询所有
	Offset *int32 `json:"offset,omitempty"`

	// 排序方式
	Order *string `json:"order,omitempty"`

	// 查询参数
	Query *string `json:"query,omitempty"`

	// 是否按名称和描述搜索
	SearchNameAndDescription *bool `json:"search_name_and_description,omitempty"`

	// 安全密级列表
	SecurityLevels *[]string `json:"security_levels,omitempty"`

	// 标签列表
	TermNames *[]string `json:"term_names,omitempty"`

	// 类型列表
	TypeNames *[]string `json:"type_names,omitempty"`
}

func (o SearchParametersExt) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchParametersExt struct{}"
	}

	return strings.Join([]string{"SearchParametersExt", string(data)}, " ")
}
