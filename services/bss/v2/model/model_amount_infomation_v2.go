package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AmountInfomationV2 struct {

	// 费用项。 具体请参见表7。
	Discounts *[]DiscountItemV2 `json:"discounts,omitempty"`

	// 现金券金额。
	FlexipurchaseCouponAmount *float64 `json:"flexipurchase_coupon_amount,omitempty"`

	// 代金券金额。
	CouponAmount *float64 `json:"coupon_amount,omitempty"`

	// 储值卡金额。
	StoredCardAmount *float64 `json:"stored_card_amount,omitempty"`

	// 手续费（仅退订订单存在）。
	CommissionAmount *float64 `json:"commission_amount,omitempty"`

	// 消费金额（仅退订订单存在）。
	ConsumedAmount *float64 `json:"consumed_amount,omitempty"`
}

func (o AmountInfomationV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AmountInfomationV2 struct{}"
	}

	return strings.Join([]string{"AmountInfomationV2", string(data)}, " ")
}
