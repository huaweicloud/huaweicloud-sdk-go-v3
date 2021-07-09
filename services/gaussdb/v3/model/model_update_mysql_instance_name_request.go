package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateMysqlInstanceNameRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlUpdateInstanceNameRequest `json:"body,omitempty"`
}

func (o UpdateMysqlInstanceNameRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateMysqlInstanceNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateMysqlInstanceNameRequest", string(data)}, " ")
}
