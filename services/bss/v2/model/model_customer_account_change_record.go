package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerAccountChangeRecord struct {

	// |参数名称：收支明细流水号| |参数约束及描述：收支明细流水号|
	AccountChangeId *string `json:"account_change_id,omitempty"`

	// |参数名称：交易详细类型| |参数约束及描述：交易详细类型|
	TradeDetailType *string `json:"trade_detail_type,omitempty"`

	// |参数名称：交易时间| |参数约束及描述：交易时间|
	TradeTime *string `json:"trade_time,omitempty"`

	// |参数名称：交易ID/订单ID| |参数约束及描述：交易ID/订单ID|
	TradeId *string `json:"trade_id,omitempty"`

	// |参数名称：变更金额| |参数约束及描述：变更金额|
	ChangeAmount *string `json:"change_amount,omitempty"`

	// |参数名称：变更后余额| |参数约束及描述：变更后余额|
	BalanceAfterChange *string `json:"balance_after_change,omitempty"`

	// |参数名称：收支类型| |参数约束及描述：收支类型|
	RevenueExpenseType *string `json:"revenue_expense_type,omitempty"`

	// |参数名称：账期| |参数约束及描述：账期|
	BillCycle *string `json:"bill_cycle,omitempty"`

	// |参数名称：交易渠道| |参数约束及描述：交易渠道|
	PaymentChannelId *string `json:"payment_channel_id,omitempty"`

	// |参数名称：交易渠道流水号| |参数约束及描述：交易渠道流水号|
	PaymentChannelNo *string `json:"payment_channel_no,omitempty"`
}

func (o CustomerAccountChangeRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerAccountChangeRecord struct{}"
	}

	return strings.Join([]string{"CustomerAccountChangeRecord", string(data)}, " ")
}
