package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type RecognizeQualificationCertificateResponse struct {
	Result         *QualificationCertificateResult `json:"result,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o RecognizeQualificationCertificateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeQualificationCertificateResponse struct{}"
	}

	return strings.Join([]string{"RecognizeQualificationCertificateResponse", string(data)}, " ")
}
