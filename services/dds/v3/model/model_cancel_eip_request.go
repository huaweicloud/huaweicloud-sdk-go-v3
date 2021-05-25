package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelEipRequest struct {
	// 节点ID。

	NodeId string `json:"node_id"`
}

func (o CancelEipRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelEipRequest struct{}"
	}

	return strings.Join([]string{"CancelEipRequest", string(data)}, " ")
}
