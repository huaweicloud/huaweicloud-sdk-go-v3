package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChargeInfo **参数说明**：设备接入实例计费信息
type ChargeInfo struct {

	// **参数说明**：实例的付费方式。 **取值范围**： - prePaid：包年/包月 - postPaid：按需计费
	ChargeMode string `json:"charge_mode"`

	// **参数说明**：订购设备接入实例的周期类型（包年、包月等）。charge_mode为prePaid时生效，且为必选值。 **取值范围**： - month：包月 - year：包年
	PeriodType *string `json:"period_type,omitempty"`

	// **参数说明**：订购设备接入实例的周期数。charge_mode为prePaid时生效，且为必选值。 **取值范围**：period_type=month（周期类型为月）时，取值为[1，9]；period_type=year（周期类型为年）时，取值为[1，3]\"
	PeriodNum *int32 `json:"period_num,omitempty"`

	// **参数说明**：创建包年/包月实例时可指定，表示是否自动续订，续订的周期和原周期相同，且续订时会自动支付。 **取值范围**： - true：自动续订 - false：默认值，不自动续订
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// **参数说明**：创建包年/包月实例时可指定，表示是否自动从客户的账户中支付，此字段不影响自动续订的支付方式。 **取值范围**：true - 自动支付，从账户余额自动扣费; false - 默认值，只提交订单不支付。[需要客户参考[\"支付包年/包月产品订单\"](https://support.huaweicloud.com/api-bpconsole/api_order_00016.html#section0)进行支付，或者在华为云官网页面使用进行支付。](tag:hws)
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o ChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargeInfo struct{}"
	}

	return strings.Join([]string{"ChargeInfo", string(data)}, " ")
}
