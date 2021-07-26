package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ApplyCertificateToHostRequest struct {
	// 证书ID

	CertificateId string `json:"certificate_id"`

	Body *ApplyCertificateToHostRequestBody `json:"body,omitempty"`
}

func (o ApplyCertificateToHostRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostRequest struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostRequest", string(data)}, " ")
}
