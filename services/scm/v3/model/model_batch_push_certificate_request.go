package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchPushCertificateRequest Request Object
type BatchPushCertificateRequest struct {

	// 证书id。
	CertificateId string `json:"certificate_id"`

	Body *BatchPushCertificateRequestBody `json:"body,omitempty"`
}

func (o BatchPushCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchPushCertificateRequest struct{}"
	}

	return strings.Join([]string{"BatchPushCertificateRequest", string(data)}, " ")
}
