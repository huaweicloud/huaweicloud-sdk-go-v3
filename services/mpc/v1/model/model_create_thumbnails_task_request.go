package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateThumbnailsTaskRequest struct {
	Body *CreateThumbReq `json:"body,omitempty"`
}

func (o CreateThumbnailsTaskRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateThumbnailsTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateThumbnailsTaskRequest", string(data)}, " ")
}
