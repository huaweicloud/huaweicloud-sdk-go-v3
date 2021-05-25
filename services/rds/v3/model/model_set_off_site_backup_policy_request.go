package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SetOffSiteBackupPolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *SetOffSiteBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o SetOffSiteBackupPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SetOffSiteBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetOffSiteBackupPolicyRequest", string(data)}, " ")
}
