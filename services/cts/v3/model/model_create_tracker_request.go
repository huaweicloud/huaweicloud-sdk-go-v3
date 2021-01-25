package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateTrackerRequest struct {
	Body *CreateTrackerRequestBody `json:"body,omitempty"`
}

func (o CreateTrackerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateTrackerRequest struct{}"
	}

	return strings.Join([]string{"CreateTrackerRequest", string(data)}, " ")
}
