package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateImageRequest struct {
	Body *CreateImageRequestBody `json:"body,omitempty"`
}

func (o CreateImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateImageRequest struct{}"
	}

	return strings.Join([]string{"CreateImageRequest", string(data)}, " ")
}
