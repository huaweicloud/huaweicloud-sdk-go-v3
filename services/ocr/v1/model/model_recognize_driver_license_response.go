package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeDriverLicenseResponse struct {
	Result         *DriverLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o RecognizeDriverLicenseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeDriverLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeDriverLicenseResponse", string(data)}, " ")
}
