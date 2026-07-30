package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapProcessInfo HTAP标准版的会话信息
type HtapProcessInfo struct {

	// **参数解释**：  会话ID。    **取值范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  会话用户名。    **取值范围**：  不涉及。
	User *string `json:"user,omitempty"`

	// **参数解释**：  会话主机。    **取值范围**：  不涉及。
	Host *string `json:"host,omitempty"`

	// **参数解释**：  会话状态。  **取值范围**：  不涉及。
	State *string `json:"state,omitempty"`

	// **参数解释**：  会话对应数据库。    **取值范围**：  不涉及。
	Database *string `json:"database,omitempty"`

	// **参数解释**：  会话执行的SQL语句。    **取值范围**：  不涉及。
	SqlStatement *string `json:"sql_statement,omitempty"`

	// **参数解释**：  会话持续时间，单位是秒。  **取值范围**：  不涉及。
	Duration *float64 `json:"duration,omitempty"`

	// **参数解释**：  会话命令类型。    **取值范围**：  不涉及。
	Command *string `json:"command,omitempty"`
}

func (o HtapProcessInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapProcessInfo struct{}"
	}

	return strings.Join([]string{"HtapProcessInfo", string(data)}, " ")
}
