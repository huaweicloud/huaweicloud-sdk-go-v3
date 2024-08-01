package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProcessesAuditLogResponse Response Object
type ShowProcessesAuditLogResponse struct {

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 日志记录集合
	ProcessAuditLogs *[]UserProcessAuditLog `json:"process_audit_logs,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ShowProcessesAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProcessesAuditLogResponse struct{}"
	}

	return strings.Join([]string{"ShowProcessesAuditLogResponse", string(data)}, " ")
}
