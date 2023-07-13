package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateRequest Request Object
type CreateCertificateRequest struct {
	Body *CreateCertificateRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateRequest", string(data)}, " ")
}
