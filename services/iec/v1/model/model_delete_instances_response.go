package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteInstancesResponse struct {
	// 边缘任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteInstancesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstancesResponse struct{}"
	}

	return strings.Join([]string{"DeleteInstancesResponse", string(data)}, " ")
}
