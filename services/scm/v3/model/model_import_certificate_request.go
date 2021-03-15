package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ImportCertificateRequest struct {
	Body *ImportCertificateRequestBody `json:"body,omitempty"`
}

func (o ImportCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportCertificateRequest struct{}"
	}

	return strings.Join([]string{"ImportCertificateRequest", string(data)}, " ")
}
