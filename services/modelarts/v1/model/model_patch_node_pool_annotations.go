package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PatchNodePoolAnnotations 更新节点池metadata的注释信息。
type PatchNodePoolAnnotations struct {

	// **参数解释**：计费模式，不指定时新创节点沿用资源池计费模式。 **取值范围**：可选值如下： - 0：按需计费 - [1：包周期计费](tag:hc)
	OsModelartsBillingMode *string `json:"os.modelarts/billing.mode,omitempty"`

	// **参数解释**：包周期订购周期，比如2。当计费模式为包周期时该参数必传。 **取值范围**：不涉及。
	OsModelartsPeriodNum *string `json:"os.modelarts/period.num,omitempty"`

	// **参数解释**：包周期订购类型。当计费模式为包周期时该参数必传。 **取值范围**：可选值如下： - 2：月 - 3：年
	OsModelartsPeriodType *string `json:"os.modelarts/period.type,omitempty"`

	// **参数解释**：是否自动续费，不指定时新创节点沿用资源池自动续费属性。 **取值范围**：可选值如下： - 0：不自动续费，默认值 - 1：自动续费
	OsModelartsAutoRenew *string `json:"os.modelarts/auto.renew,omitempty"`

	// **参数解释**：用户在运营平台选择的折扣信息。 **取值范围**：不涉及。
	OsModelartsPromotionInfo *string `json:"os.modelarts/promotion.info,omitempty"`

	// **参数解释**：订购订单支付完成后跳转的url地址。 **取值范围**：不涉及。
	OsModelartsServiceConsoleUrl *string `json:"os.modelarts/service.console.url,omitempty"`

	// **参数解释**：订单id，包周期资源创建或者计费模式变更的时候该参数必需。 **取值范围**：不涉及。
	OsModelartsOrderId *string `json:"os.modelarts/order.id,omitempty"`
}

func (o PatchNodePoolAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PatchNodePoolAnnotations struct{}"
	}

	return strings.Join([]string{"PatchNodePoolAnnotations", string(data)}, " ")
}
