package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HraRuleCreateRequest HRA规则
type HraRuleCreateRequest struct {

	// **参数解释：** 扩缩容类型。 **取值范围：** - SIMULATOR_ALGO：模拟器算法扩缩容类型。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	ScalerType *string `json:"scaler_type,omitempty"`

	// **参数解释：** SLO配置参数信息。 **取值范围：** 不涉及。
	SloInfo []SloInfo `json:"slo_info"`

	// **参数解释：** 指标信息。 **取值范围：** 不涉及。
	Metrics []Metrics `json:"metrics"`

	// **参数解释：** 角色扩缩策略（不会进行实质扩缩，因此该配置值无效）。 **取值范围：** 1~128。
	RoleReplica *[]RoleReplica `json:"role_replica,omitempty"`
}

func (o HraRuleCreateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HraRuleCreateRequest struct{}"
	}

	return strings.Join([]string{"HraRuleCreateRequest", string(data)}, " ")
}
