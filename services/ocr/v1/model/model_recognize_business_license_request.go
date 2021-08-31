package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeBusinessLicenseRequest struct {
	Body *BusinessLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeBusinessLicenseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeBusinessLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessLicenseRequest", string(data)}, " ")
}
