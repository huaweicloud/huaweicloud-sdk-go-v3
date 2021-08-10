package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunImageTaggingRequest struct {
	Body *ImageTaggingReq `json:"body,omitempty"`
}

func (o RunImageTaggingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunImageTaggingRequest struct{}"
	}

	return strings.Join([]string{"RunImageTaggingRequest", string(data)}, " ")
}
