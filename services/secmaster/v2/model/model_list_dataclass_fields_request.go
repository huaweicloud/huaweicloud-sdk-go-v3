package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataclassFieldsRequest Request Object
type ListDataclassFieldsRequest struct {

	// 内容类型
	ContentType string `json:"content-type"`

	// 工作空间id
	WorkspaceId string `json:"workspace_id"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 数据量
	Limit *int32 `json:"limit,omitempty"`

	// 名称查询
	Name *string `json:"name,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 数据类id
	DataclassId string `json:"dataclass_id"`

	// 字段分类
	FieldCategory *string `json:"field_category,omitempty"`

	// 是否展示在分类映射外的其他地方
	Mapping *bool `json:"mapping,omitempty"`
}

func (o ListDataclassFieldsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataclassFieldsRequest struct{}"
	}

	return strings.Join([]string{"ListDataclassFieldsRequest", string(data)}, " ")
}
