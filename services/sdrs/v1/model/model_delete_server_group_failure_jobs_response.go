package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteServerGroupFailureJobsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteServerGroupFailureJobsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteServerGroupFailureJobsResponse struct{}"
	}

	return strings.Join([]string{"DeleteServerGroupFailureJobsResponse", string(data)}, " ")
}
