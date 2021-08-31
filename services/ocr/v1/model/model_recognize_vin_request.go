package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeVinRequest struct {
	Body *VinRequestBody `json:"body,omitempty"`
}

func (o RecognizeVinRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeVinRequest struct{}"
	}

	return strings.Join([]string{"RecognizeVinRequest", string(data)}, " ")
}
