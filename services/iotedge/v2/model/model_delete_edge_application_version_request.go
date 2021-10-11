package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEdgeApplicationVersionRequest struct {
	// 应用版本,应用内版本唯一。

	EdgeAppId string `json:"edge_app_id"`
	// 应用版本ID，应用版本唯一。

	Version string `json:"version"`
}

func (o DeleteEdgeApplicationVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeApplicationVersionRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeApplicationVersionRequest", string(data)}, " ")
}
