package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunImageTaggingResponse struct {
	Result         *ImageTaggingResponseResult `json:"result,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o RunImageTaggingResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunImageTaggingResponse struct{}"
	}

	return strings.Join([]string{"RunImageTaggingResponse", string(data)}, " ")
}
