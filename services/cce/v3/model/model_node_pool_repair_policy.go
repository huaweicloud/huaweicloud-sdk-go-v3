package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolRepairPolicy 节点故障自愈配置
type NodePoolRepairPolicy struct {

	// **参数解释**： 系统与 K8s 组件异常时是否启用policy中配置的自愈策略。 **约束限制**： 不涉及 **取值范围**： - false：使用基础自愈策略 - true：使用policy中配置的自愈策略  **默认取值**： false
	Enable *bool `json:"enable,omitempty"`

	// **参数解释**： 节点自愈的恢复策略 **约束限制**： - 当 enable 为 true 时，此字段必填。 - 当 enable 为 false 时，此字段无效，用户填写任意值均不会生效，系统使用基础自愈策略。  **取值范围**： - restartNode：系统与 K8s 组件异常时允许通过重启节点自愈  **默认取值**： 不涉及
	Policy *string `json:"policy,omitempty"`
}

func (o NodePoolRepairPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolRepairPolicy struct{}"
	}

	return strings.Join([]string{"NodePoolRepairPolicy", string(data)}, " ")
}
