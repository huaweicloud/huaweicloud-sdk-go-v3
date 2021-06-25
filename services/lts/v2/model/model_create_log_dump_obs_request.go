package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateLogDumpObsRequest struct {
	Body *CreateLogDumpObsRequestBody `json:"body,omitempty"`
}

func (o CreateLogDumpObsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateLogDumpObsRequest struct{}"
	}

	return strings.Join([]string{"CreateLogDumpObsRequest", string(data)}, " ")
}
