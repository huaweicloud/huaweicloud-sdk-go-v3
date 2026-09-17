package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeCheckStatus **参数解释：** 节点限制检查状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type NodeCheckStatus struct {

	// **参数解释：** 状态 **约束限制：** 不涉及 **取值范围：** - Init：初始化 - Running：运行中 - Success：成功 - Failed：失败  **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释：** 节点检查状态 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	NodeStageStatus *[]NodeStageStatus `json:"nodeStageStatus,omitempty"`
}

func (o NodeCheckStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeCheckStatus struct{}"
	}

	return strings.Join([]string{"NodeCheckStatus", string(data)}, " ")
}
