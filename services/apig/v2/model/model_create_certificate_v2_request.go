package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCertificateV2Request struct {
	Body *CertificateForm `json:"body,omitempty"`
}

func (o CreateCertificateV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateV2Request struct{}"
	}

	return strings.Join([]string{"CreateCertificateV2Request", string(data)}, " ")
}
