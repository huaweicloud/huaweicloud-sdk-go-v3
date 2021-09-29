package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteGaussMySqlInstanceRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`
}

func (o DeleteGaussMySqlInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlInstanceRequest struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlInstanceRequest", string(data)}, " ")
}
