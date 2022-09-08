package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiscountItemV2 struct {

	// 折扣类型： 200：促销产品折扣300：促销折扣券301：促销代金券302：促销现金券500：代理订购指定折扣501：代理订购指定减免502：代理订购指定一口价600：折扣返利合同601：渠道框架合同602：专款专用合同603：线下直签合同604：电销授权合同605：商务合同折扣606：渠道商务合同折扣607：合作伙伴授权折扣609：订单调价折扣610：免单金额700：促销折扣800：充值帐户折扣900：产品本身折扣901：基准价一口价的折扣
	DiscountType *string `json:"discount_type,omitempty"`

	// 折扣金额。
	DiscountAmount *float64 `json:"discount_amount,omitempty"`
}

func (o DiscountItemV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscountItemV2 struct{}"
	}

	return strings.Join([]string{"DiscountItemV2", string(data)}, " ")
}
