package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAllServerGroupFailureJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAllServerGroupFailureJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAllServerGroupFailureJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAllServerGroupFailureJobsResponse", string(data)}, " ")
}
