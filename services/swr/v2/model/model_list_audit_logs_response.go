package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditLogsResponse Response Object
type ListAuditLogsResponse struct {

	// 审计日志列表
	AuditLogs *[]AuditLog `json:"audit_logs,omitempty"`

	// 审计日志总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuditLogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditLogsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditLogsResponse", string(data)}, " ")
}
