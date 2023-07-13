package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateQuotaResponse Response Object
type ShowCertificateQuotaResponse struct {
	Quotas         *Quotas `json:"quotas,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCertificateQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowCertificateQuotaResponse", string(data)}, " ")
}
