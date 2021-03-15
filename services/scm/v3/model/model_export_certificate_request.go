package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ExportCertificateRequest struct {
	CertificateId string `json:"certificate_id"`
}

func (o ExportCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportCertificateRequest struct{}"
	}

	return strings.Join([]string{"ExportCertificateRequest", string(data)}, " ")
}
