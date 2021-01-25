package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateResetTracksTaskRequest struct {
	Body *CreateResetTracksReq `json:"body,omitempty"`
}

func (o CreateResetTracksTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateResetTracksTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateResetTracksTaskRequest", string(data)}, " ")
}
