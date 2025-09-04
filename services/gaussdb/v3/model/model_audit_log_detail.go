package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditLogDetail 审计日志信息。
type AuditLogDetail struct {

	// **参数解释**：  审计日志ID。  **取值范围**：  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  审计日志文件名。  **取值范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  审计日志大小，单位：KB。  **取值范围**：  不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**：  审计日志开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。  **取值范围**：  不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**：  审计日志结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始，Z指时区偏移量，例如北京时间偏移显示为+0800。  **取值范围**：  不涉及。
	EndTime *string `json:"end_time,omitempty"`
}

func (o AuditLogDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditLogDetail struct{}"
	}

	return strings.Join([]string{"AuditLogDetail", string(data)}, " ")
}
