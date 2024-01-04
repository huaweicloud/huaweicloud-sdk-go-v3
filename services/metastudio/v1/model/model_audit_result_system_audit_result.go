package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AuditResultSystemAuditResult 系统审核结果。
type AuditResultSystemAuditResult struct {

	// 操作时间。
	AuditTime *int64 `json:"audit_time,omitempty"`

	// 错误列表。
	Errors *[]AuditResultSystemAuditResultErrors `json:"errors,omitempty"`
}

func (o AuditResultSystemAuditResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AuditResultSystemAuditResult struct{}"
	}

	return strings.Join([]string{"AuditResultSystemAuditResult", string(data)}, " ")
}
