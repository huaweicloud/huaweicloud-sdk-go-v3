package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunImageModerationResponse struct {
	Result         *ImageDetectionResultBody `json:"result,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o RunImageModerationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunImageModerationResponse struct{}"
	}

	return strings.Join([]string{"RunImageModerationResponse", string(data)}, " ")
}
