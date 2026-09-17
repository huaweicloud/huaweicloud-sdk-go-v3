package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PreCheckItemStatus **参数解释：** 检查项状态信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
type PreCheckItemStatus struct {

	// **参数解释：** 检查项名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 检查项类型 **约束限制：** 不涉及 **取值范围：** - Exception：异常类，需要用户解决 - Risk：风险类，用户确认后可选择跳过  **默认取值：** 不涉及
	Kind *string `json:"kind,omitempty"`

	// **参数解释：** 检查项分组 **约束限制：** 不涉及 **取值范围：** - LimitCheck：集群限制检查 - MasterCheck：控制节点检查 - NodeCheck：用户节点检查 - AddonCheck：插件检查 - ExecuteException：检查流程错误  **默认取值：** 不涉及
	Group *string `json:"group,omitempty"`

	// **参数解释：** 检查项风险级别 **约束限制：** 不涉及 **取值范围：** - Info：提示级别 - Warning：风险级别 - Fatal：严重级别  **默认取值：** 不涉及
	Level *string `json:"level,omitempty"`

	// **参数解释：** 状态 **约束限制：** 不涉及 **取值范围：** - Init：初始化 - Running：运行中 - Success：成功 - Failed：失败  **默认取值：** 不涉及
	Phase *string `json:"phase,omitempty"`

	// **参数解释：** 提示信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Message *string `json:"message,omitempty"`

	RiskSource *RiskSource `json:"riskSource,omitempty"`

	// **参数解释：** 错误码集合 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ErrorCodes *[]string `json:"errorCodes,omitempty"`
}

func (o PreCheckItemStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreCheckItemStatus struct{}"
	}

	return strings.Join([]string{"PreCheckItemStatus", string(data)}, " ")
}
