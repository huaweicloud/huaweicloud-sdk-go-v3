package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteInstancesRequest struct {
	Body *DeleteInstancesRequestBody `json:"body,omitempty"`
}

func (o DeleteInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteInstancesRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstancesRequest", string(data)}, " ")
}
