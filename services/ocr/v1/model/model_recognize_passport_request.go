package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizePassportRequest struct {
	Body *PassportRequestBody `json:"body,omitempty"`
}

func (o RecognizePassportRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizePassportRequest struct{}"
	}

	return strings.Join([]string{"RecognizePassportRequest", string(data)}, " ")
}
