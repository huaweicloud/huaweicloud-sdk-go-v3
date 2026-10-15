package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProgressInfo 迁移进度信息体
type ProgressInfo struct {

	// **参数解释**：  完成进度，以百分比形式表示，如\"94%\"、\"100%\"。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Completed *string `json:"completed,omitempty"`

	// **参数解释**：  该子项迁移的预计剩余时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	RemainingTime *string `json:"remaining_time,omitempty"`
}

func (o ProgressInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgressInfo struct{}"
	}

	return strings.Join([]string{"ProgressInfo", string(data)}, " ")
}
