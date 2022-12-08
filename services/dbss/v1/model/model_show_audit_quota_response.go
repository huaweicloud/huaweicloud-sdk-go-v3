package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAuditQuotaResponse struct {

	// 实例当前剩余配额。
	AuditQuota *int64 `json:"audit_quota,omitempty"`

	// Cpu当前剩余配额。
	Cpu *int64 `json:"cpu,omitempty"`

	// 项目Id。
	ProjectId *string `json:"project_id,omitempty"`

	// 配额。
	Quota *int64 `json:"quota,omitempty"`

	// 内存当前剩余配额
	Ram            *int64 `json:"ram,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAuditQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAuditQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditQuotaResponse", string(data)}, " ")
}
