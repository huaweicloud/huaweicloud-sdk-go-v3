package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowEnterpriseProjectQuotaResponse struct {
	Quotas         *QuotasDetail `json:"quotas,omitempty" xml:"quotas"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowEnterpriseProjectQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEnterpriseProjectQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowEnterpriseProjectQuotaResponse", string(data)}, " ")
}
