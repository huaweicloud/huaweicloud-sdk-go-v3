package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RunTextModerationResponse struct {
	Result         *TextDetectionResponseResult `json:"result,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o RunTextModerationResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RunTextModerationResponse struct{}"
	}

	return strings.Join([]string{"RunTextModerationResponse", string(data)}, " ")
}
