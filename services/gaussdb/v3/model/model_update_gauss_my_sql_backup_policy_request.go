package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateGaussMySqlBackupPolicyRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlUpdateBackupPolicyRequest `json:"body,omitempty"`
}

func (o UpdateGaussMySqlBackupPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateGaussMySqlBackupPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateGaussMySqlBackupPolicyRequest", string(data)}, " ")
}
