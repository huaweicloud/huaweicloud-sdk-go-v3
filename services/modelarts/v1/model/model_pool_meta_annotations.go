package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMetaAnnotations 资源池metadata的注释信息。
type PoolMetaAnnotations struct {

	// **参数解释**：资源池的描述信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsDescription *string `json:"os.modelarts/description,omitempty"`

	// **参数解释**：计费模式。 **约束限制**：不涉及。 **取值范围**：可选值如下： - 0：按需计费 - 1：包周期计费 **默认取值**：不涉及。
	OsModelartsBillingMode *string `json:"os.modelarts/billing.mode,omitempty"`

	// **参数解释**：包周期资源池的订购周期。 **约束限制**：和os.modelarts/period.type字段配合使用。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPeriodNum *string `json:"os.modelarts/period.num,omitempty"`

	// **参数解释**：包周期资源池的订购类型。 **约束限制**：和os.modelarts/period.num字段配合使用。 **取值范围**：可选值如下： - 2：包月。 - 3：包年。 **默认取值**：不涉及。
	OsModelartsPeriodType *string `json:"os.modelarts/period.type,omitempty"`

	// **参数解释**：包周期资源池的自动续费类型。 **约束限制**：不涉及。 **取值范围**：可选值如下： - 0：不自动续费。 - 1：自动续费。 **默认取值**：0。
	OsModelartsAutoRenew *string `json:"os.modelarts/auto.renew,omitempty"`

	// **参数解释**：包周期资源池购买时选择的折扣信息。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPromotionInfo *string `json:"os.modelarts/promotion.info,omitempty"`

	// **参数解释**：购买包周期资源在订单支付完成后跳转地址。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsServiceConsoleUrl *string `json:"os.modelarts/service.console.url,omitempty"`

	// **参数解释**：包周期资源池购买时传递的订单ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsOrderId *string `json:"os.modelarts/order.id,omitempty"`

	// **参数解释**：包周期资源池中资源规格对应的资源ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsFlavorResourceIds *string `json:"os.modelarts/flavor.resource.ids,omitempty"`

	// **参数解释**：资源池上的资源标签。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsTmsTags *string `json:"os.modelarts/tms.tags,omitempty"`

	// **参数解释**：资源池调度队列的策略，用于定义任务调度的规则。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPoolSchedulerQueueStrategy *string `json:"os.modelarts.pool/scheduler.queue.strategy,omitempty"`

	// **参数解释**：资源池包含的逻辑子池的数量。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPoolSubpoolsCount *string `json:"os.modelarts.pool/subpools.count,omitempty"`

	// **参数解释**：资源池的租户账号 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsTenantDomainName *string `json:"os.modelarts/tenant.domain.name,omitempty"`

	// **参数解释**：训练外部依赖标识 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	OsModelartsPoolScopeExternalDependencyTrain *string `json:"os.modelarts.pool/scope.external.dependency.train,omitempty"`
}

func (o PoolMetaAnnotations) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMetaAnnotations struct{}"
	}

	return strings.Join([]string{"PoolMetaAnnotations", string(data)}, " ")
}
