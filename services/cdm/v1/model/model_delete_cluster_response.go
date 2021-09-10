package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteClusterResponse struct {
	// 作业ID

	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteClusterResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteClusterResponse struct{}"
	}

	return strings.Join([]string{"DeleteClusterResponse", string(data)}, " ")
}
