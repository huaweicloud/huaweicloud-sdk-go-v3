package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartnerConsumptionQuotaRequest Request Object
type ShowPartnerConsumptionQuotaRequest struct {
}

func (o ShowPartnerConsumptionQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartnerConsumptionQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowPartnerConsumptionQuotaRequest", string(data)}, " ")
}
