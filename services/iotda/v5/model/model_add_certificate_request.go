package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AddCertificateRequest struct {
	InstanceId *string               `json:"Instance-Id,omitempty"`
	Body       *CreateCertificateDto `json:"body,omitempty"`
}

func (o AddCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AddCertificateRequest struct{}"
	}

	return strings.Join([]string{"AddCertificateRequest", string(data)}, " ")
}
