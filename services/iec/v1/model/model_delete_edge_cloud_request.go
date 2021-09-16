package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteEdgeCloudRequest struct {
	// 边缘业务ID。

	EdgecloudId string `json:"edgecloud_id"`
}

func (o DeleteEdgeCloudRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeCloudRequest struct{}"
	}

	return strings.Join([]string{"DeleteEdgeCloudRequest", string(data)}, " ")
}
