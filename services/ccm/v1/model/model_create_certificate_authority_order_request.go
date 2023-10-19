package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateAuthorityOrderRequest Request Object
type CreateCertificateAuthorityOrderRequest struct {
	Body *CreateCertificateAuthorityOrderRequestBody `json:"body,omitempty"`
}

func (o CreateCertificateAuthorityOrderRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityOrderRequest struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityOrderRequest", string(data)}, " ")
}
