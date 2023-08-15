package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditlogsResponse Response Object
type ListAuditlogsResponse struct {

	// 总记录数。
	TotalRecord *int32 `json:"total_record,omitempty"`

	// 审计日志具体信息。
	AuditLogs      *[]ListAuditlogsResult `json:"audit_logs,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListAuditlogsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditlogsResponse struct{}"
	}

	return strings.Join([]string{"ListAuditlogsResponse", string(data)}, " ")
}
