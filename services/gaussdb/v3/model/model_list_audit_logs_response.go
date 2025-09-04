package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditLogsResponse Response Object
type ListAuditLogsResponse struct {

	// **参数解释**：  记录列表。  **取值范围**：  不涉及。
	AuditLogs *[]AuditLogDetail `json:"audit_logs,omitempty"`

	// **参数解释**：  总记录数。  **取值范围**：  不小于0。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditLogsResponse", string(data)}, " ")
}
