package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditResultAdminAuditResult 管理员审核结果。
type AuditResultAdminAuditResult struct {

	// 审核信息。
	Message *string `json:"message,omitempty"`

	// 附件名称。
	AttachmentName *string `json:"attachment_name,omitempty"`

	// 附件下载地址。
	AttachmentUrl *string `json:"attachment_url,omitempty"`

	// 操作时间。
	AuditTime *int64 `json:"audit_time,omitempty"`
}

func (o AuditResultAdminAuditResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditResultAdminAuditResult struct{}"
	}

	return strings.Join([]string{"AuditResultAdminAuditResult", string(data)}, " ")
}
