package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IntelligentKillSessionHistory struct {

	// **参数解释**：  智能Kill会话动作任务ID。  **取值范围**：  不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**：  智能Kill会话动作起始时间。  **取值范围**：  开始执行智能Kill会话动作时刻的秒级时间戳。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**：  智能Kill会话动作结束时间。  **取值范围**：  结束执行智能Kill会话动作时刻的秒级时间戳。
	EndTime *int64 `json:"end_time,omitempty"`

	// **参数解释**：  智能Kill会话历史记录下载链接。  **取值范围**：  不涉及。
	DownloadLink *string `json:"download_link,omitempty"`
}

func (o IntelligentKillSessionHistory) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IntelligentKillSessionHistory struct{}"
	}

	return strings.Join([]string{"IntelligentKillSessionHistory", string(data)}, " ")
}
