package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteCertificateRequest struct {
	// 证书id。

	CertificateId string `json:"certificate_id"`
}

func (o DeleteCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteCertificateRequest struct{}"
	}

	return strings.Join([]string{"DeleteCertificateRequest", string(data)}, " ")
}
