package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ImportCertificateResponse struct {
	// 证书id。
	CertificateId  *string `json:"certificate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportCertificateResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ImportCertificateResponse", string(data)}, " ")
}
