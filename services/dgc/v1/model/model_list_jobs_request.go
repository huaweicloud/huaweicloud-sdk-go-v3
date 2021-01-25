package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListJobsRequest struct {
}

func (o ListJobsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListJobsRequest struct{}"
	}

	return strings.Join([]string{"ListJobsRequest", string(data)}, " ")
}
