package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPlaybookAuditLogsResponse Response Object
type ListPlaybookAuditLogsResponse struct {

	// 总条数
	Count *int32 `json:"count,omitempty"`

	// 审计日志列表信息
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
