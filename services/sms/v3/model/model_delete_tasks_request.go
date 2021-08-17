package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteTasksRequest struct {
	Body *DeleteTasksReq `json:"body,omitempty"`
}

func (o DeleteTasksRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTasksRequest struct{}"
	}

	return strings.Join([]string{"DeleteTasksRequest", string(data)}, " ")
}
