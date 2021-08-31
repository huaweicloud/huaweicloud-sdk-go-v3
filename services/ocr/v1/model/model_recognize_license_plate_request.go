package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeLicensePlateRequest struct {
	Body *LicensePlateRequestBody `json:"body,omitempty"`
}

func (o RecognizeLicensePlateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeLicensePlateRequest struct{}"
	}

	return strings.Join([]string{"RecognizeLicensePlateRequest", string(data)}, " ")
}
