package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelineTemplatesQuery struct {

	// 创建模板时，用户选择的语言
	Language *string `json:"language,omitempty"`

	// 是否系统模板
	IsSystem *bool `json:"is_system,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0，默认为0
	Offset *int64 `json:"offset,omitempty"`

	// 每次查询的条目数量，默认为10。
	Limit *int64 `json:"limit,omitempty"`

	// 用于排序的字段，非必选。取值为：name，create_time
	SortKey *string `json:"sort_key,omitempty"`

	// 排序类型，非必选。asc按排序字段升序，desc按排序字段降序
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListPipelineTemplatesQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineTemplatesQuery struct{}"
	}

	return strings.Join([]string{"ListPipelineTemplatesQuery", string(data)}, " ")
}
