package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResizeInstance struct {
	Flavor *Flavor `json:"flavor"`

	// **参数说明**：修改包年/包月实例的规格时可指定，表示是否自动从客户的账户中支付，此字段不影响自动续订的支付方式。 **取值范围**：true - 自动支付，从账户余额自动扣费; false - 默认值，只提交订单不支付。[需要客户参考[\"支付包年/包月产品订单\"](https://support.huaweicloud.com/api-bpconsole/api_order_00016.html#section0)进行支付，或者在华为云官网页面使用进行支付。](tag:hws)
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// **参数说明**：是否延时变更设备实例的计费信息。约束：如需延时变更，需要先设置实例的变更时间窗。 **取值范围**： - true：延迟变更，规格变更任务将在指定的变更时间窗内执行。 - false：立即变更，规格变更任务将立即执行。
	Delay *bool `json:"delay,omitempty"`
}

func (o ResizeInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResizeInstance struct{}"
	}

	return strings.Join([]string{"ResizeInstance", string(data)}, " ")
}
