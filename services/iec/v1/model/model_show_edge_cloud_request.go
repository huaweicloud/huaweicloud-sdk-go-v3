package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEdgeCloudRequest struct {
	// 边缘业务ID。

	EdgecloudId string `json:"edgecloud_id"`
}

func (o ShowEdgeCloudRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEdgeCloudRequest struct{}"
	}

	return strings.Join([]string{"ShowEdgeCloudRequest", string(data)}, " ")
}
