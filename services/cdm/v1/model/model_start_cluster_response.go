package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type StartClusterResponse struct {
	// 作业ID

	JobId          *[]string `json:"jobId,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o StartClusterResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartClusterResponse struct{}"
	}

	return strings.Join([]string{"StartClusterResponse", string(data)}, " ")
}
