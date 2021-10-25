package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RecognizeQualificationCertificateRequest struct {
	Body *QualificationCertificateRequestBody `json:"body,omitempty"`
}

func (o RecognizeQualificationCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RecognizeQualificationCertificateRequest struct{}"
	}

	return strings.Join([]string{"RecognizeQualificationCertificateRequest", string(data)}, " ")
}
