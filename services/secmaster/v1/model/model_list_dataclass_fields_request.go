package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataclassFieldsRequest Request Object
type ListDataclassFieldsRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// 名称查询
	Name *string `json:"name,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// **参数解释：** 数据类id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
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
