package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteIntelligentKillSessionResponse Response Object
type ExecuteIntelligentKillSessionResponse struct {

	// **参数解释**：  智能Kill会话任务ID。  **取值范围**：  不涉及。
	TaskId         *string `json:"task_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteIntelligentKillSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteIntelligentKillSessionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteIntelligentKillSessionResponse", string(data)}, " ")
}
