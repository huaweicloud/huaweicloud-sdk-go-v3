package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateMysqlBackupPolicyResponse struct {
	// 状态信息

	Status *string `json:"status,omitempty"`
	// 实例ID

	InstanceId *string `json:"instance_id,omitempty"`
	// 实例名称

	InstanceName   *string `json:"instance_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateMysqlBackupPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMysqlBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateMysqlBackupPolicyResponse", string(data)}, " ")
}
