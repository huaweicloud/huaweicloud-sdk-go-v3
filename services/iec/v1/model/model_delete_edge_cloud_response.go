package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteEdgeCloudResponse struct {
	// 边缘任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteEdgeCloudResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteEdgeCloudResponse struct{}"
	}

	return strings.Join([]string{"DeleteEdgeCloudResponse", string(data)}, " ")
}
