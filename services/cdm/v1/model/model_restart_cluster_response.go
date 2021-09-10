package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RestartClusterResponse struct {
	// 作业ID

	JobId          *string `json:"jobId,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RestartClusterResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestartClusterResponse struct{}"
	}

	return strings.Join([]string{"RestartClusterResponse", string(data)}, " ")
}
