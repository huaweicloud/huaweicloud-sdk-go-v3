package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCertificateRequest struct {
	// 企业项目id

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateCertificateRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequest", string(data)}, " ")
}
