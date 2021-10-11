package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEdgeApplicationVersionRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`
	// 应用版本,应用内版本唯一。

	Version string `json:"version"`
}

func (o ShowEdgeApplicationVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEdgeApplicationVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeApplicationVersionRequest", string(data)}, " ")
}
