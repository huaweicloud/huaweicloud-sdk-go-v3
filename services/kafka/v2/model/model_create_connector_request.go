package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateConnectorRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *CreateConnectorReq `json:"body,omitempty"`
}

func (o CreateConnectorRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateConnectorRequest struct{}"
	}

	return strings.Join([]string{"CreateConnectorRequest", string(data)}, " ")
}
