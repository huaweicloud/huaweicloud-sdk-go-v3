package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckCertificateRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	CertificateId string `json:"certificate_id"`

	ActionId string `json:"action_id"`

	Body *VerifyCertificateDto `json:"body,omitempty"`
}

func (o CheckCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckCertificateRequest struct{}"
	}

	return strings.Join([]string{"CheckCertificateRequest", string(data)}, " ")
}
