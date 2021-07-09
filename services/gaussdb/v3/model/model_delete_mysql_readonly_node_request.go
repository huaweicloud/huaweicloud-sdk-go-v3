package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteMysqlReadonlyNodeRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID，严格匹配UUID规则。

	InstanceId string `json:"instance_id"`
	// 节点ID，严格匹配UUID规则。

	NodeId string `json:"node_id"`
}

func (o DeleteMysqlReadonlyNodeRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteMysqlReadonlyNodeRequest struct{}"
	}

	return strings.Join([]string{"DeleteMysqlReadonlyNodeRequest", string(data)}, " ")
}
