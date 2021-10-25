package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAllServerGroupFailureJobsRequest struct {
}

func (o DeleteAllServerGroupFailureJobsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAllServerGroupFailureJobsRequest struct{}"
	}

	return strings.Join([]string{"DeleteAllServerGroupFailureJobsRequest", string(data)}, " ")
}
