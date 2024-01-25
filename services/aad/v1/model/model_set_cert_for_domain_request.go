package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetCertForDomainRequest Request Object
type SetCertForDomainRequest struct {
	Body *CertificateBody `json:"body,omitempty"`
}

func (o SetCertForDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetCertForDomainRequest struct{}"
	}

	return strings.Join([]string{"SetCertForDomainRequest", string(data)}, " ")
}
