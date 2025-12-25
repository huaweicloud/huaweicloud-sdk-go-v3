package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentActionsRequest Request Object
type ListComponentActionsRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// 插件id
	ComponentId string `json:"component_id"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`

	// 是否启用
	Enabled bool `json:"enabled"`
}

func (o ListComponentActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentActionsRequest struct{}"
	}

	return strings.Join([]string{"ListComponentActionsRequest", string(data)}, " ")
}
