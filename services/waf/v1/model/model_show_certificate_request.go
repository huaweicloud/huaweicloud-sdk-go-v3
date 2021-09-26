package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCertificateRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 证书ID

	CertificateId string `json:"certificate_id"`
}

func (o ShowCertificateRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCertificateRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateRequest", string(data)}, " ")
}
