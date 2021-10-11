package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateEdgeApplicationVersionRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`
	// 应用版本,应用内版本唯一。

	Version string `json:"version"`

	Body *UpdateEdgeAppVersionDto `json:"body,omitempty"`
}

func (o UpdateEdgeApplicationVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEdgeApplicationVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdateEdgeApplicationVersionRequest", string(data)}, " ")
}
