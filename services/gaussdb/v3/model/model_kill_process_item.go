package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KillProcessItem **参数解释**：  Kill会话任务条件，条件各项之间是“与”的逻辑关系。
type KillProcessItem struct {

	// **参数解释**：  任务ID。  **取值范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例ID。  **取值范围**：  只能由英文字母、数字组成，后缀为in07，长度为36个字符。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  用户名。  **取值范围**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  主机。  **取值范围**：  不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：  数据库名。  **取值范围**：  不涉及。
	Db *string `json:"db,omitempty"`

	// **参数解释**：  命令。  **取值范围**：  不涉及。
	Command *string `json:"command,omitempty"`

	// **参数解释**：  SQL语句。  **取值范围**：  不涉及。
	Info *string `json:"info,omitempty"`

	// **参数解释**：  持续时间，单位为秒。持续时间大于等于该值的会话满足Kill条件。  **取值范围**：  由数字组成，表示的范围应为1-2147483647。
	Time *string `json:"time,omitempty"`

	// **参数解释**：  任务结束时间戳。  **取值范围**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**：  执行方式。  **取值范围**：  - auto：定时关闭，任务持续时长等于end_time参数值。 - manual：手动关闭。
	Mode *string `json:"mode,omitempty"`

	// **参数解释**：  任务状态。  **取值范围**：  - 0：已结束。 - 1：运行中。
	Status *string `json:"status,omitempty"`
}

func (o KillProcessItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KillProcessItem struct{}"
	}

	return strings.Join([]string{"KillProcessItem", string(data)}, " ")
}
