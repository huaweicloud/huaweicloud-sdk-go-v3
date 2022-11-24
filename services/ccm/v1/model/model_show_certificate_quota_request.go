package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCertificateQuotaRequest struct {
}

func (o ShowCertificateQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateQuotaRequest", string(data)}, " ")
}
