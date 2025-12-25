package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionsRequest Request Object
type ListConnectionsRequest struct {

	// **参数解释：** 内容类型 - application/json;charset=UTF-8    普通API请求的类型  **约束限制：** 不涉及 **取值范围：** - application/json;charset=UTF-8  **默认取值：** 不涉及
	ContentType string `json:"content-type"`

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// 连接名称
	Name *string `json:"name,omitempty"`

	// 插件名称
	ComponentName *string `json:"component_name,omitempty"`

	// 创建人
	CreatorName *string `json:"creator_name,omitempty"`

	// 修改人
	ModifierName *string `json:"modifier_name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 创建起始时间
	CreateStartTime *string `json:"create_start_time,omitempty"`

	// 创建结束时间
	CreateEndTime *string `json:"create_end_time,omitempty"`

	// 修改起始时间
	UpdateStartTime *string `json:"update_start_time,omitempty"`

	// 修改结束时间
	UpdateEndTime *string `json:"update_end_time,omitempty"`

	// 是否用于应急策略的操作连接
	IsDefenseType *bool `json:"is_defense_type,omitempty"`
}

func (o ListConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionsRequest", string(data)}, " ")
}
