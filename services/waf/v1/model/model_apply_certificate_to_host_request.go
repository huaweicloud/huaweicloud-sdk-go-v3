package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ApplyCertificateToHostRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 证书ID

	CertificateId string `json:"certificate_id"`

	Body *ApplyCertificateToHostRequestBody `json:"body,omitempty"`
}

func (o ApplyCertificateToHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostRequest struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostRequest", string(data)}, " ")
}
