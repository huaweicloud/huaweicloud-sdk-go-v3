package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAuditlogPolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowAuditlogPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowAuditlogPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAuditlogPolicyRequest", string(data)}, " ")
}
