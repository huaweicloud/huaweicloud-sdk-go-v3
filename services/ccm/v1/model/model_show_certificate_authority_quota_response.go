package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateAuthorityQuotaResponse Response Object
type ShowCertificateAuthorityQuotaResponse struct {
	Quotas         *Quotas `json:"quotas,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCertificateAuthorityQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityQuotaResponse", string(data)}, " ")
}
