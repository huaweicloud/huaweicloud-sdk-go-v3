package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateAuthorityQuotaRequest Request Object
type ShowCertificateAuthorityQuotaRequest struct {
}

func (o ShowCertificateAuthorityQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityQuotaRequest", string(data)}, " ")
}
