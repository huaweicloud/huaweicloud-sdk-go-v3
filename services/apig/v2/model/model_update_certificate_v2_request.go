package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateCertificateV2Request struct {

	// 证书的编号
	CertificateId string `json:"certificate_id"`

	Body *CertificateForm `json:"body,omitempty"`
}

func (o UpdateCertificateV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateV2Request struct{}"
	}

	return strings.Join([]string{"UpdateCertificateV2Request", string(data)}, " ")
}
