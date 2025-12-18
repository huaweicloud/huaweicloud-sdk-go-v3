package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScaleUpBillingConfigOverride 节点池扩容时覆盖节点的默认计费模式配置
type ScaleUpBillingConfigOverride struct {

	// **参数解释**： 节点计费类型 **约束限制**： 选填参数，不填表示使用节点池默认计费配置 **取值范围**： - 0：按需 - 1：包周期 **默认取值**： 不涉及
	BillingMode *int32 `json:"billingMode,omitempty"`

	ExtendParam *ScaleUpExtendParam `json:"extendParam,omitempty"`
}

func (o ScaleUpBillingConfigOverride) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScaleUpBillingConfigOverride struct{}"
	}

	return strings.Join([]string{"ScaleUpBillingConfigOverride", string(data)}, " ")
}
