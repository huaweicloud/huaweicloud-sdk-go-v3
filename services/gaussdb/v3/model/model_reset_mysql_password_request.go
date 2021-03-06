package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ResetMysqlPasswordRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlResetPasswordRequest `json:"body,omitempty"`
}

func (o ResetMysqlPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ResetMysqlPasswordRequest struct{}"
	}

	return strings.Join([]string{"ResetMysqlPasswordRequest", string(data)}, " ")
}
