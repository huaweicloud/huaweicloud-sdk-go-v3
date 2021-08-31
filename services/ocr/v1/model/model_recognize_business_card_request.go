package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeBusinessCardRequest struct {
	Body *BusinessCardRequestBody `json:"body,omitempty"`
}

func (o RecognizeBusinessCardRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeBusinessCardRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessCardRequest", string(data)}, " ")
}
