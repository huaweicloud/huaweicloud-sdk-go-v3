package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateAuthorityRequest Request Object
type CreateCertificateAuthorityRequest struct {
	Body *CreateCertificateAuthorityRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityRequest", string(data)}, " ")
}
