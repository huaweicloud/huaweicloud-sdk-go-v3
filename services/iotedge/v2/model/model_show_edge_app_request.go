package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEdgeAppRequest struct {
	// 应用ID，应用唯一。

	EdgeAppId string `json:"edge_app_id"`
}

func (o ShowEdgeAppRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEdgeAppRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeAppRequest", string(data)}, " ")
}
