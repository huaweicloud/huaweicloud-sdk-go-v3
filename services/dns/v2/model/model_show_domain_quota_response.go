package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowDomainQuotaResponse struct {
	Quotas         *[]DomainQuotaResponseQuotas `json:"quotas,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ShowDomainQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainQuotaResponse", string(data)}, " ")
}
