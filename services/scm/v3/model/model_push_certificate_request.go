package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type PushCertificateRequest struct {
	// 证书id。

	CertificateId string `json:"certificate_id"`

	Body *PushCertificateRequestBody `json:"body,omitempty"`
}

func (o PushCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushCertificateRequest struct{}"
	}

	return strings.Join([]string{"PushCertificateRequest", string(data)}, " ")
}
