package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMysqlBackupPolicyRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`
}

func (o ShowMysqlBackupPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowMysqlBackupPolicyRequest", string(data)}, " ")
}
