package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAuditlogsResponse struct {
	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`
	// 审计日志具体信息。
	AuditLogs      *[]ListAuditlogsResult `json:"audit_logs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAuditlogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAuditlogsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditlogsResponse", string(data)}, " ")
}
