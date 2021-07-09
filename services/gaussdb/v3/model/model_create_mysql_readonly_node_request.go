package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMysqlReadonlyNodeRequest struct {
	// 语言。

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`

	Body *MysqlCreateReadonlyNodeRequest `json:"body,omitempty"`
}

func (o CreateMysqlReadonlyNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlReadonlyNodeRequest struct{}"
	}

	return strings.Join([]string{"CreateMysqlReadonlyNodeRequest", string(data)}, " ")
}
