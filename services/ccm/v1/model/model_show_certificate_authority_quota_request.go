package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCertificateAuthorityQuotaRequest struct {
}

func (o ShowCertificateAuthorityQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityQuotaRequest", string(data)}, " ")
}
