package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RunImageModerationRequest struct {
	Body *ImageDetectionReq `json:"body,omitempty"`
}

func (o RunImageModerationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunImageModerationRequest struct{}"
	}

	return strings.Join([]string{"RunImageModerationRequest", string(data)}, " ")
}
