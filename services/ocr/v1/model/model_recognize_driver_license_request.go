package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeDriverLicenseRequest struct {
	Body *DriverLicenseRequestBody `json:"body,omitempty"`
}

func (o RecognizeDriverLicenseRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeDriverLicenseRequest struct{}"
	}

	return strings.Join([]string{"RecognizeDriverLicenseRequest", string(data)}, " ")
}
