package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KillProcessTaskReq **参数解释**：  Kill会话任务条件，条件各项之间是“与”的逻辑关系。  **约束限制**：  支持同时运行最多5个条件Kill任务。
type KillProcessTaskReq struct {

	// **参数解释**：  用户名。  **约束限制**：  不涉及。  **取值范围**：  1到32个字符之间，由字母、数字、中划线或下划线组成，不能包含其他特殊字符。  **默认取值**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  主机。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：  数据库名。  **约束限制**：  不涉及。  **取值范围**：  1到64个字符之间，由字母、数字、中划线或下划线组成，不能包含其他特殊字符。   **默认取值**：  不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：  命令。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：  SQL语句。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Info *string `json:"info,omitempty"`

	// **参数解释**：  持续时间，单位为秒。持续时间大于等于该值的会话满足Kill条件。  **约束限制**：  必须是整数。  **取值范围**：  1-2147483647。  **默认取值**：  不涉及。
	Time string `json:"time"`

	// **参数解释**：  任务持续时间，单位为秒。  **约束限制**：  执行方式为auto时必填。  **取值范围**：  10-31535999。  **默认取值**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  执行方式。  **约束限制**：  不涉及。  **取值范围**：  - auto：定时关闭，任务持续时长等于end_time参数值。 - manual：手动关闭。  **默认取值**：  不涉及。
	Mode string `json:"mode"`
}

func (o KillProcessTaskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KillProcessTaskReq struct{}"
	}

	return strings.Join([]string{"KillProcessTaskReq", string(data)}, " ")
}
