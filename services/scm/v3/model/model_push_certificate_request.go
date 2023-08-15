package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushCertificateRequest Request Object
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
