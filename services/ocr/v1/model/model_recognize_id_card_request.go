package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeIdCardRequest struct {
	Body *IdCardRequestBody `json:"body,omitempty"`
}

func (o RecognizeIdCardRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeIdCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeIdCardRequest", string(data)}, " ")
}
