package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChargeInfoOption 计费类型信息，支持包年包月和按需计费，默认为按需计费。
type ChargeInfoOption struct {

	// **参数解释：** 计费模式。 **约束限制：** 不涉及。 **取值范围：** - prePaid：预付费，即包年/包月。 - postPaid：后付费，即按需付费。 **默认取值：** 不涉及。
	ChargeMode string `json:"charge_mode"`

	// **参数解释：** 订购周期类型。 **约束限制：** “charge_mode”为“prePaid”时生效，且为必选值。 **取值范围：** - month：包月。 - year：包年。 **默认取值：** 不涉及。
	PeriodType *string `json:"period_type,omitempty"`

	// **参数解释：** 指定订购的时间。 **约束限制：** “charge_mode”为“prePaid”时生效，且为必选值。 **取值范围：** - 当“period_type”为“month”时，取值为1~9。 - 当“period_type”为“year”时，取值为1~3 **默认取值：** 不涉及。
	PeriodNum *string `json:"period_num,omitempty"`

	// **参数解释：** 创建包周期实例时可指定，表示是否自动续订，续订的周期和原周期相同，且续订时会自动支付。 **约束限制：** 不涉及。 **取值范围：** - true，表示自动续订。 - false，表示不自动续订，默认为该方式。 **默认取值：** false。
	IsAutoRenew *string `json:"is_auto_renew,omitempty"`

	// **参数解释：** 创建包周期实例时可指定，表示是否自动从账户中支付，该字段不影响自动续订的支付方式。 **约束限制：** 不涉及。 **取值范围：** - true，表示自动从账户中支付 - false，表示手动从账户中支付，默认为该支付方式 **默认取值：** false。
	IsAutoPay *string `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfoOption struct{}"
	}

	return strings.Join([]string{"ChargeInfoOption", string(data)}, " ")
}
