package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetDdlLogPolicyResponse Response Object
type SetDdlLogPolicyResponse struct {

	// **参数解释**：  DDL下载日志列表。  **取值范围**：  不涉及。
	DdlLogs *[]DdlLogInfo `json:"ddl_logs,omitempty"`

	// **参数解释**：  总条数。  **取值范围**：  不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  日志保留天数。  **取值范围**：  不涉及。
	KeeyDays *int32 `json:"keey_days,omitempty"`

	// **参数解释**：  DDL日志下载开关状态。  **取值范围**：  - ON，开启。 - OFF，关闭。
	SwitchStatus   *string `json:"switch_status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetDdlLogPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetDdlLogPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetDdlLogPolicyResponse", string(data)}, " ")
}
