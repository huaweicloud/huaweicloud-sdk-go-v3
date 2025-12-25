package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataclassRequest Request Object
type ListDataclassRequest struct {

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

	// 业务编码
	BusinessCode *string `json:"business_code,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`
}

func (o ListDataclassRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataclassRequest struct{}"
	}

	return strings.Join([]string{"ListDataclassRequest", string(data)}, " ")
}
