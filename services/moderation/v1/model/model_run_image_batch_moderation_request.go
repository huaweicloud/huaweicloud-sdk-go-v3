package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunImageBatchModerationRequest struct {
	Body *ImageBatchModerationReq `json:"body,omitempty"`
}

func (o RunImageBatchModerationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunImageBatchModerationRequest struct{}"
	}

	return strings.Join([]string{"RunImageBatchModerationRequest", string(data)}, " ")
}
