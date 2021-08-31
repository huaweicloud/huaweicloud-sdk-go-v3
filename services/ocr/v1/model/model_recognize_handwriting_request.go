package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeHandwritingRequest struct {
	Body *HandwritingRequestBody `json:"body,omitempty"`
}

func (o RecognizeHandwritingRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeHandwritingRequest struct{}"
	}

	return strings.Join([]string{"RecognizeHandwritingRequest", string(data)}, " ")
}
