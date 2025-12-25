package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentConfigurationRequest Request Object
type ListComponentConfigurationRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 组件ID
	ComponentId string `json:"component_id"`

	// 每页数量
	Limit *int64 `json:"limit,omitempty"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// 按照属性排序。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序顺序，支持 `ASC` 或 `DESC`。
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListComponentConfigurationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentConfigurationRequest struct{}"
	}

	return strings.Join([]string{"ListComponentConfigurationRequest", string(data)}, " ")
}
