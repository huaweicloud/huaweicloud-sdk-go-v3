package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// ShowCustomerExpenditureQuotaResponse Response Object
type ShowCustomerExpenditureQuotaResponse struct {

	// |参数名称：消费配额| |参数约束及描述：当前授予额度，消费配额。|
	GrantingAmount *decimal.Decimal `json:"granting_amount,omitempty"`

	// |参数名称：已使用消费配额| |参数约束及描述：已使用消费配额。|
	GrantingAmountUsed *decimal.Decimal `json:"granting_amount_used,omitempty"`

	// |参数名称：已使用消费配额占比| |参数约束及描述：已使用消费配额占比。|
	PercentageGrantingAmountUsed *decimal.Decimal `json:"percentage_granting_amount_used,omitempty"`

	// |参数名称：金额的度量单位。| |参数约束及描述：1：元|
	MeasureId *int32 `json:"measure_id,omitempty"`

	// |参数名称：币种。| |参数约束及描述：货币单位，USD：美元。|
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCustomerExpenditureQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerExpenditureQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerExpenditureQuotaResponse", string(data)}, " ")
}
