package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPlaybookAuditLogsResponse struct {

	// tatal count
	Count *int32 `json:"count,omitempty"`

	// list of informations of Audit Log Info
	AuditLogs *[]AuditLogInfo `json:"audit_logs,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPlaybookAuditLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPlaybookAuditLogsResponse struct{}"
	}

	return strings.Join([]string{"ListPlaybookAuditLogsResponse", string(data)}, " ")
}
