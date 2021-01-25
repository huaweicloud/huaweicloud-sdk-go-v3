package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListJobsResponse struct {
	Total          *int32     `json:"total,omitempty"`
	Jobs           *[]JobInfo `json:"jobs,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListJobsResponse struct{}"
	}

	return strings.Join([]string{"ListJobsResponse", string(data)}, " ")
}
