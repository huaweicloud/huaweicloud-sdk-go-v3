package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAuditQuotaNewResponse Response Object
type ShowAuditQuotaNewResponse struct {

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 审计实例剩余配额。
	AuditQuota *int64 `json:"audit_quota,omitempty"`

	// CPU剩余配额。
	Cpu *int64 `json:"cpu,omitempty"`

	// 内存剩余配额。
	Ram            *int64 `json:"ram,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAuditQuotaNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditQuotaNewResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditQuotaNewResponse", string(data)}, " ")
}
