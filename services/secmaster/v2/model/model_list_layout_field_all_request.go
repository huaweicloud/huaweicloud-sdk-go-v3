package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLayoutFieldAllRequest Request Object
type ListLayoutFieldAllRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释：** 工作空间id。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 名称查询
	Name *string `json:"name,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 数据类业务编码
	BusinessCode *string `json:"business_code,omitempty"`

	// 布局id
	LayoutId *string `json:"layout_id,omitempty"`
}

func (o ListLayoutFieldAllRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLayoutFieldAllRequest struct{}"
	}

	return strings.Join([]string{"ListLayoutFieldAllRequest", string(data)}, " ")
}
