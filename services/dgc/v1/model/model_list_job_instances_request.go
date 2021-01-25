package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListJobInstancesRequest struct {
}

func (o ListJobInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListJobInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListJobInstancesRequest", string(data)}, " ")
}
