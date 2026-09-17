package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeTaskStatus **参数解释：** 升级任务状态信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type UpgradeTaskStatus struct {

	// **参数解释：** 升级任务状态 **约束限制：** 不涉及 **取值范围：** - Init：初始化 - Queuing：等待 - Running：运行中 - Pause：暂停 - Success：成功 - Failed：失败  **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释：** 升级任务进度 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Progress *string `json:"progress,omitempty"`

	// **参数解释：** 升级任务结束时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CompletionTime *string `json:"completionTime,omitempty"`
}

func (o UpgradeTaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeTaskStatus struct{}"
	}

	return strings.Join([]string{"UpgradeTaskStatus", string(data)}, " ")
}
