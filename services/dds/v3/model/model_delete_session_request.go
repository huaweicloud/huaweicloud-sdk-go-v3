package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteSessionRequest struct {
	// 节点ID。允许查询的节点如下： 集群下面的 mongos节点以及 副本集、单节点实例下面的所有节点。

	NodeId string `json:"node_id"`

	Body *DeleteSessionRequestBody `json:"body,omitempty"`
}

func (o DeleteSessionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteSessionRequest struct{}"
	}

	return strings.Join([]string{"DeleteSessionRequest", string(data)}, " ")
}
