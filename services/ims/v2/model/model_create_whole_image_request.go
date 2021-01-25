package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateWholeImageRequest struct {
	Body *CreateWholeImageRequestBody `json:"body,omitempty"`
}

func (o CreateWholeImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWholeImageRequest struct{}"
	}

	return strings.Join([]string{"CreateWholeImageRequest", string(data)}, " ")
}
