package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerAccountChangeRecordsRequest Request Object
type ListCustomerAccountChangeRecordsRequest struct {

	// |参数名称：账户类型。BALANCE_TYPE_DEBIT：现金账户BALANCE_TYPE_CREDIT：信用账户| |参数的约束及描述：必填|
	BalanceType string `json:"balance_type"`

	// |参数名称：收支类型。REVENUE：收入EXPENSE：支出| |参数的约束及描述：非必填|
	RevenueExpenseType *string `json:"revenue_expense_type,omitempty"`

	// |参数名称：交易类型。RECHARGE：充值DEDEUCT：消费REFUND：退款RFROZEN：冻结TRANS：转账，余额和保证金互换（老商务模式，当前已无保证金账户）ADJUST：调账（华为核销等）BEADJUST：经销商拨款BERETRIEVE：经销商回收BEUNBIND：解绑/关联模式切换导致的回收EXPIRED：过期清零BONUSCONVERT：奖励金转换（老商务模式，当前已无奖励金账户）TRADE_MODE_TRANSFER：交易模式变更DEPOSIT：保证金| |参数的约束及描述：非必填|
	TradeType *string `json:"trade_type,omitempty"`

	// |参数名称：查询收支明细的开始日期| |参数的约束及描述：非必填|
	TradeTimeBegin *string `json:"trade_time_begin,omitempty"`

	// |参数名称：查询收支明细的结束日期| |参数的约束及描述：非必填|
	TradeTimeEnd *string `json:"trade_time_end,omitempty"`

	// |参数名称：交易ID/订单ID| |参数的约束及描述：非必填|
	TradeId *string `json:"trade_id,omitempty"`

	// |参数名称：交易渠道。可以一次查询多个，用逗号分隔。1：支付宝2：银行转账4：支付宝/网银5：微信支付6：提现7：激励返点10：交易模式变更11：调账317：银联（统一支付平台）319：Huawei Pay320：华为支付| |参数的约束及描述：非必填|
	PaymentChannelId *string `json:"payment_channel_id,omitempty"`

	// |参数名称：交易渠道流水号| |参数的约束及描述：非必填|
	PaymentChannelNo *string `json:"payment_channel_no,omitempty"`

	// |参数名称：偏移量| |参数的约束及描述：非必填|
	Offset *int32 `json:"offset,omitempty"`

	// |参数名称：每页的显示条数| |参数的约束及描述：非必填|
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomerAccountChangeRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerAccountChangeRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerAccountChangeRecordsRequest", string(data)}, " ")
}
