package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowOffSiteBackupPolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o ShowOffSiteBackupPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowOffSiteBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowOffSiteBackupPolicyRequest", string(data)}, " ")
}
