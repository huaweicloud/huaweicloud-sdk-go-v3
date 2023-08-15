package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQuotasResponse Response Object
type ShowQuotasResponse struct {
	Quotas *QuotaNoLimit `json:"quotas,omitempty"`

	// 站点配额信息，暂不包括中心站点
	SiteQuotas     *[]SiteQuotaNoLimit `json:"site_quotas,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ShowQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotasResponse", string(data)}, " ")
}
