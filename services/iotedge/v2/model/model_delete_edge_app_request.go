package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEdgeAppRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`
}

func (o DeleteEdgeAppRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeAppRequest", string(data)}, " ")
}
