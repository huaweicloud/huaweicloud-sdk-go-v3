package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProcessDetail struct {

	// **参数解释**  会话ID。  **取值范围**：  不涉及。
	Id *int64 `json:"id,omitempty"`

	// **参数解释**：  用户名。  **取值范围**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：    数据库名。    **取值范围**：    不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：  执行状态。  **取值范围**：  不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**：  命令。  **取值范围**：  不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：  SQL语句。  **取值范围**：  不涉及。
	SqlInfo *string `json:"sql_info,omitempty"`

	// **参数解释**：  状态持续时间，单位是秒。  **取值范围**：  不涉及。
	StateDuration *int64 `json:"state_duration,omitempty"`
}

func (o ProcessDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessDetail struct{}"
	}

	return strings.Join([]string{"ProcessDetail", string(data)}, " ")
}
