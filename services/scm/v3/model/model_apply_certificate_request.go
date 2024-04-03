package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyCertificateRequest Request Object
type ApplyCertificateRequest struct {

	// 证书id。
	CertificateId string `json:"certificate_id"`

	Body *ApplyCertificateRequestBody `json:"body,omitempty"`
}

func (o ApplyCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateRequest struct{}"
	}

	return strings.Join([]string{"ApplyCertificateRequest", string(data)}, " ")
}
