package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KillProcessHistoryItem struct {

	// **参数解释**：  自动Kill会话任务ID。  **取值范围**：  不涉及。
	TaskId *string `json:"task_id,omitempty"`

	// **参数解释**：  会话ID。  **取值范围**：  不涉及。
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**：  Kill会话的时间戳。  **取值范围**：  不涉及。
	KillTime *string `json:"kill_time,omitempty"`

	// **参数解释**：  Kill会话方式。  **取值范围**：  - once：一键Kill会话。 - condition：条件Kill会话。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**：  用户。  **取值范围**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  主机。  **取值范围**：  不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：  数据库名。  **取值范围**：  不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：  命令。  **取值范围**：  不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：  会话持续时间，单位是秒。  **取值范围**：  不涉及。
	Time *string `json:"time,omitempty"`

	// **参数解释**：  SQL语句。  **取值范围**：  不涉及。
	Info *string `json:"info,omitempty"`

	// **参数解释**：  执行状态。  **取值范围**：  不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**：  实例ID。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  节点ID。  **取值范围**：  只能由英文字母、数字组成，后缀为no07，长度为36个字符。
	NodeId *string `json:"node_id,omitempty"`
}

func (o KillProcessHistoryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KillProcessHistoryItem struct{}"
	}

	return strings.Join([]string{"KillProcessHistoryItem", string(data)}, " ")
}
