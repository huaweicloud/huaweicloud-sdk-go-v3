package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type PushCertificateRequest struct {
	CertificateId string                      `json:"certificate_id"`
	Body          *PushCertificateRequestBody `json:"body,omitempty"`
}

func (o PushCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PushCertificateRequest struct{}"
	}

	return strings.Join([]string{"PushCertificateRequest", string(data)}, " ")
}
