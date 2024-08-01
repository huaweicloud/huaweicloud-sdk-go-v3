package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartnerConsumptionQuotaResponse Response Object
type ShowPartnerConsumptionQuotaResponse struct {

	// |参数名称：总消费配额| |参数约束及描述：非必填|
	ConsumptionQuota *string `json:"consumption_quota,omitempty"`

	// |参数名称：账户余额| |参数约束及描述：非必填|
	Amount *string `json:"amount,omitempty"`

	// |参数名称：代金券余额| |参数约束及描述：非必填|
	CouponAmount *string `json:"coupon_amount,omitempty"`

	// |参数名称：应还金额| |参数约束及描述：非必填|
	DueAmount *string `json:"due_amount,omitempty"`

	// |参数名称：未出账预估金额| |参数约束及描述：非必填|
	UnbilledAmount *string `json:"unbilled_amount,omitempty"`

	// |参数名称：已使用消费配额| |参数约束及描述：非必填，used_consumption_quota = due_amount + unbilled_amount - amount - coupon_amount|
	UsedConsumptionQuota *string `json:"used_consumption_quota,omitempty"`

	// |参数名称：金额单位| |参数约束及描述：金额单位，1：元|
	MeasureId *int32 `json:"measure_id,omitempty"`

	// |参数名称：货币单位| |参数约束及描述：货币单位 |
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPartnerConsumptionQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartnerConsumptionQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowPartnerConsumptionQuotaResponse", string(data)}, " ")
}
