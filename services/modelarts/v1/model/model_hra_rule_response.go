package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HraRuleResponse HRA规则
type HraRuleResponse struct {

	// **参数解释：** 规则ID，在创建HRA策略时即可在返回体中获取，也可通过查询推理单元配比检测信息获取当前用户拥有的HRA策略，其中id字段即为规则ID。 **约束限制：** 不涉及。 **取值范围：** 规则ID。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 规则名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 规则是否禁用。 **取值范围：** - true：禁用。 - false：不禁用。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Disable *bool `json:"disable,omitempty"`

	// **参数解释：** 扩缩容类型。 **取值范围：** - SIMULATOR_ALGO：模拟器算法扩缩容类型。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	ScalerType *string `json:"scaler_type,omitempty"`

	// **参数解释：** HRA规则状态。 **取值范围：** - CREATING：创建。 - CONFIG_SUCCESS：配置HRA策略成功。 - EXECUTE_SUCCESS：执行HRA策略成功。 - DELETED：删除。 - FAILED：失败。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	RuleStatus *string `json:"rule_status,omitempty"`

	// **参数解释：** SLO配置参数信息。 **取值范围：** 不涉及。
	SloInfo *[]SloInfo `json:"slo_info,omitempty"`

	// **参数解释：** 指标信息。 **取值范围：** 不涉及。
	Metrics *[]Metrics `json:"metrics,omitempty"`

	// **参数解释：** 角色扩缩策略（不会进行实质扩缩，因此该配置值无效）。
	RoleReplica *[]RoleReplica `json:"role_replica,omitempty"`
}

func (o HraRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HraRuleResponse struct{}"
	}

	return strings.Join([]string{"HraRuleResponse", string(data)}, " ")
}
