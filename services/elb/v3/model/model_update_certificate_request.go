package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCertificateRequest struct {
	// 证书ID。

	CertificateId string `json:"certificate_id"`

	Body *UpdateCertificateRequestBody `json:"body,omitempty"`
}

func (o UpdateCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateRequest struct{}"
	}

	return strings.Join([]string{"UpdateCertificateRequest", string(data)}, " ")
}
