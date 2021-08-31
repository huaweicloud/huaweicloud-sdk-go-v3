package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeWebImageRequest struct {
	Body *WebImageRequestBody `json:"body,omitempty"`
}

func (o RecognizeWebImageRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeWebImageRequest struct{}"
	}

	return strings.Join([]string{"RecognizeWebImageRequest", string(data)}, " ")
}
