package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type DiscountDetailInfo struct {

	// 折扣类型。 500：代理订购指定折扣 501：代理订购指定减免 502：代理订购指定一口价 600：合同折扣返利 （商履折扣） 601：渠道框架合同折扣 602：专款专用合同折扣（特殊商务合同折扣） 603：线下直签合同折扣 604：电销授权合同折扣 605：商务合同折扣 606：渠道商务合同折扣 607：合作伙伴授权折扣 608：严选商品折扣 610：免单金额 700：促销折扣 （促销，只有包年/包月场景） 800：赠送奖励金
	PromotionType *string `json:"promotion_type,omitempty"`

	// 折扣金额。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 折扣类型对应的标识，可为合同ID或商务ID。
	PromotionId *string `json:"promotion_id,omitempty"`

	// 金额单位，1:元 3：分，默认3
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o DiscountDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscountDetailInfo struct{}"
	}

	return strings.Join([]string{"DiscountDetailInfo", string(data)}, " ")
}
