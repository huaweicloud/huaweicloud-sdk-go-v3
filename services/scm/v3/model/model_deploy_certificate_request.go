package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeployCertificateRequest Request Object
type DeployCertificateRequest struct {

	// 证书id。
	CertificateId string `json:"certificate_id"`

	Body *DeployCertificateRequestBody `json:"body,omitempty"`
}

func (o DeployCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeployCertificateRequest struct{}"
	}

	return strings.Join([]string{"DeployCertificateRequest", string(data)}, " ")
}
