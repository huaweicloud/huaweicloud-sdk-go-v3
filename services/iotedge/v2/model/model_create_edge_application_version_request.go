package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateEdgeApplicationVersionRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`

	Body *CreateEdgeApplicationVersionDto `json:"body,omitempty"`
}

func (o CreateEdgeApplicationVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationVersionRequest struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationVersionRequest", string(data)}, " ")
}
