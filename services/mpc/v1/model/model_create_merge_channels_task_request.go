package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMergeChannelsTaskRequest struct {
	Body *CreateMergeChannelsReq `json:"body,omitempty"`
}

func (o CreateMergeChannelsTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMergeChannelsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateMergeChannelsTaskRequest", string(data)}, " ")
}
