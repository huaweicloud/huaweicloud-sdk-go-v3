package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentsRequest Request Object
type ListComponentsRequest struct {

	// **参数解释**: 工作空间ID **取值范围**: 不涉及
	WorkspaceId string `json:"workspace_id"`

	// **参数解释：** 偏移量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 数据量 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListComponentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentsRequest struct{}"
	}

	return strings.Join([]string{"ListComponentsRequest", string(data)}, " ")
}
