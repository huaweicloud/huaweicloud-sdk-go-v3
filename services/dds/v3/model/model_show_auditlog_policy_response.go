package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowAuditlogPolicyResponse struct {
	// 审计日志保存天数，审计日志策略关闭时为0。

	KeepDays *int32 `json:"keep_days,omitempty"`
	// 审计范围。

	AuditScope *string `json:"audit_scope,omitempty"`
	// 审计类型。

	AuditTypes     *[]string `json:"audit_types,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAuditlogPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAuditlogPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowAuditlogPolicyResponse", string(data)}, " ")
}
