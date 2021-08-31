package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeBusinessLicenseResponse struct {
	Result         *BusinessLicenseResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o RecognizeBusinessLicenseResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeBusinessLicenseResponse struct{}"
	}

	return strings.Join([]string{"RecognizeBusinessLicenseResponse", string(data)}, " ")
}
