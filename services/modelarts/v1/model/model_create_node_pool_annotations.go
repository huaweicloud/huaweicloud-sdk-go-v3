package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNodePoolAnnotations 创建节点池metadata的注释信息。
type CreateNodePoolAnnotations struct {

	// **参数解释**：计费模式。 **约束限制**：不涉及。 **取值范围**：可选值如下： - 0：按需计费 - 1：包周期计费 **默认取值**：不涉及。
	OsModelartsBillingMode *string `json:"os.modelarts/billing.mode,omitempty"`

	// **参数解释**：包周期订购周期，比如2。当计费模式为包周期时该参数必传。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPeriodNum *string `json:"os.modelarts/period.num,omitempty"`

	// **参数解释**：包周期订购类型。当计费模式为包周期时该参数必传 **约束限制**：不涉及。 **取值范围**：可选值如下： - 2：月 - 3：年 - 4：小时 **默认取值**：不涉及。
	OsModelartsPeriodType *string `json:"os.modelarts/period.type,omitempty"`

	// **参数解释**：是否自动续费。 **约束限制**：不涉及。 **取值范围**：可选值如下： - 0：不自动续费 - 1：自动续费 **默认取值**：0。
	OsModelartsAutoRenew *string `json:"os.modelarts/auto.renew,omitempty"`

	// **参数解释**：用户在运营平台选择的折扣信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPromotionInfo *string `json:"os.modelarts/promotion.info,omitempty"`

	// **参数解释**：订购订单支付完成后跳转的url地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsServiceConsoleUrl *string `json:"os.modelarts/service.console.url,omitempty"`

	// **参数解释**：订单id，包周期资源创建或者计费模式变更的时候该参数必需。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsOrderId *string `json:"os.modelarts/order.id,omitempty"`
}

func (o CreateNodePoolAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNodePoolAnnotations struct{}"
	}

	return strings.Join([]string{"CreateNodePoolAnnotations", string(data)}, " ")
}
