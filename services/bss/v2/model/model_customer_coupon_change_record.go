package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerCouponChangeRecord struct {

	// |参数名称：优惠券ID| |参数约束及描述：优惠券ID，范围限制:0-64|
	CouponId *string `json:"coupon_id,omitempty"`

	// |参数名称：交易详细类型| |参数约束及描述：交易详细类型，范围限制:0-128， SOURCE_OPERAIION_DEDEUCT：消费(包年/包月) SOURCE_OPERAIION_USAGE：消费(按需) SOURCE_OPERAIION_SAVINGS_PLANS：消费(节省计划) SOURCE_OPERAIION_WRITEOFF：消费(欠费还款) SOURCE_OPERAIION_EXPIRECLEAR：过期清零 SOURCE_OPERAIION_UNSUBSCRIBE：退券(退订) SOURCE_OPERAIION_UNFROZEN：退券(流程中) SOURCE_OPERAIION_RFROZEN：退券(流程中) SOURCE_OPERAIION_PRIZE：激活 SOURCE_OPERAIION_BILLREFUND：调账(华为核销) SOURCE_OPERATION_COUPONCANCEL：优惠券回收 SOURCE_OPERATION_DEPOSIT_FROZEN：保证金冻结 SOURCE_OPERATION_DEPOSIT_UNFROZEN：保证金解冻 SOURCE_OPERATION_RETRIEVE：回收(企业回收) SOURCE_OPERATION_TRANSFER：划拨(企业划拨)|
	TradeDetailType *string `json:"trade_detail_type,omitempty"`

	// |参数名称：交易时间| |参数约束及描述：交易时间，UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ|
	TradeTime *string `json:"trade_time,omitempty"`

	// |参数名称：交易ID/订单ID| |参数约束及描述：交易ID/订单ID，范围限制：0-128|
	TradeId *string `json:"trade_id,omitempty"`

	// |参数名称：变更金额| |参数约束及描述：变更金额|
	ChangeAmount *string `json:"change_amount,omitempty"`

	// |参数名称：变更后余额| |参数约束及描述：变更后余额|
	BalanceAfterChange *string `json:"balance_after_change,omitempty"`

	// |参数名称：收支类型| |参数约束及描述：收支类型。REVENUE：收入 EXPENSE：支出|
	RevenueExpenseType *string `json:"revenue_expense_type,omitempty"`

	// |参数名称：账期| |参数约束及描述：账期|
	BillCycle *string `json:"bill_cycle,omitempty"`

	// |参数名称：账号名称| |参数约束及描述：账号名称，范围限制：0-128|
	AccountName *string `json:"account_name,omitempty"`

	// |参数名称：云服务类型名称| |参数约束及描述：产品的云服务名称，范围限制：0-200|
	CloudServiceTypeName *string `json:"cloud_service_type_name,omitempty"`

	// |参数名称：资源类型名称| |参数约束及描述：产品的资源类型名称，范围限制：0-200|
	ResourceTypeName *string `json:"resource_type_name,omitempty"`
}

func (o CustomerCouponChangeRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerCouponChangeRecord struct{}"
	}

	return strings.Join([]string{"CustomerCouponChangeRecord", string(data)}, " ")
}
