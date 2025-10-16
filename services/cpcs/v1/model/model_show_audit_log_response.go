package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditLogResponse Response Object
type ShowAuditLogResponse struct {

	// 日志总数
	Total *int32 `json:"total,omitempty"`

	// 日志列表
	AuditRecords   *[]ShowAuditLogResponseBodyAuditRecords `json:"audit_records,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ShowAuditLogResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditLogResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditLogResponse", string(data)}, " ")
}
