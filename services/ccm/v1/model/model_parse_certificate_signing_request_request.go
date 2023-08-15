package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ParseCertificateSigningRequestRequest Request Object
type ParseCertificateSigningRequestRequest struct {
	Body *ParseCertificateSigningRequestRequestBody `json:"body,omitempty"`
}

func (o ParseCertificateSigningRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParseCertificateSigningRequestRequest struct{}"
	}

	return strings.Join([]string{"ParseCertificateSigningRequestRequest", string(data)}, " ")
}
