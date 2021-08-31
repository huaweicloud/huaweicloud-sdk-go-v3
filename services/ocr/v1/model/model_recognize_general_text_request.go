package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeGeneralTextRequest struct {
	Body *GeneralTextRequestBody `json:"body,omitempty"`
}

func (o RecognizeGeneralTextRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeGeneralTextRequest struct{}"
	}

	return strings.Join([]string{"RecognizeGeneralTextRequest", string(data)}, " ")
}
