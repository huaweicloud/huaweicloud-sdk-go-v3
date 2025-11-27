package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerCouponChangeRecordsRequest Request Object
type ListCustomerCouponChangeRecordsRequest struct {

	// |参数名称：账户类型| |参数的约束及描述：该参数必填，BALANCE_TYPE_COUPON：代金券账户|
	BalanceType string `json:"balance_type"`

	// |参数名称：收支类型| |参数的约束及描述：该参数非必填，REVENUE：收入 EXPENSE：支出。此参数不携带时，不作为筛选条件；此参数携带值不允许为空、空串，有枚举值校验。|
	RevenueExpenseType *string `json:"revenue_expense_type,omitempty"`

	// |参数名称：交易类型| |参数的约束及描述：该参数非必填，ADJUST：激活 DEDEUCT：消费 REFUND：退券 RFROZEN：冻结 EXPIRED：过期清零 COUPONADJUST：划拨 COUPONCANCEL：回收。此参数不携带时，不作为筛选条件；此参数携带值不允许为空、空串，有枚举值校验。|
	TradeType *string `json:"trade_type,omitempty"`

	// |参数名称：交易ID/订单ID| |参数的约束及描述：该参数非必填，范围限制：0-128。此参数不携带或携带值为空时，不作为筛选条件。|
	TradeId *string `json:"trade_id,omitempty"`

	// |参数名称：查询收支明细的开始日期| |参数的约束及描述：该参数非必填，东八区时间，格式为YYYY-MM-DD，此参数不携带、携带值为空时，默认值为一年前的当天日期；此参数不允许携带值为空串，有参数校验。|
	TradeTimeBegin *string `json:"trade_time_begin,omitempty"`

	// |参数名称：查询收支明细的结束日期| |参数的约束及描述：该参数非必填，东八区时间，格式为YYYY-MM-DD，此参数不携带、携带值为空时，默认值为当前日期；此参数不允许携带值为空串，有参数校验。|
	TradeTimeEnd *string `json:"trade_time_end,omitempty"`

	// |参数名称：优惠券ID。| |参数的约束及描述：该参数非必填，范围限制：0-64。此参数不携带或携带值为空时，不作为筛选条件。|
	CouponId *string `json:"coupon_id,omitempty"`

	// |参数名称：偏移量| |参数的约束及描述：该参数非必填，范围限制：0-2147483647，默认值为0。此参数不携带或携带值为空时，默认传参为0。|
	Offset *int32 `json:"offset,omitempty"`

	// |参数名称：每次查询的数量| |参数的约束及描述：该参数非必填，范围限制：1-100，默认值为10。此参数不携带或携带值为空时，默认传参为10。|
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListCustomerCouponChangeRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerCouponChangeRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerCouponChangeRecordsRequest", string(data)}, " ")
}
