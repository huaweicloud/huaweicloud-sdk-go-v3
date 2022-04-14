package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowSubCustomerBudgetResponse struct {
	// 初始预算金额。

	BudgetAmount *float64 `json:"budget_amount,omitempty"`
	// 已经使用的预算。该预算存在一定的时延和误差。

	UsedAmount *float64 `json:"used_amount,omitempty"`
	// 金额单位。 1：元

	MeasureId *int32 `json:"measure_id,omitempty"`
	// 币种。 USD：美金

	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSubCustomerBudgetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSubCustomerBudgetResponse struct{}"
	}

	return strings.Join([]string{"ShowSubCustomerBudgetResponse", string(data)}, " ")
}
