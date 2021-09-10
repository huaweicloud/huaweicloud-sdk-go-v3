package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StopClusterResponse struct {
	// 作业ID

	JobId          *[]string `json:"jobId,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o StopClusterResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StopClusterResponse struct{}"
	}

	return strings.Join([]string{"StopClusterResponse", string(data)}, " ")
}
