package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HraRuleUpdateRequest HRA规则
type HraRuleUpdateRequest struct {

	// 规则ID，在[创建HRA策略](CreateInferHra.xml)时即可在返回体中获取，也可通过[获取推理单元配比检测信息](ShowInferHra.xml)获取当前用户拥有的HRA策略，其中id字段即为规则ID。
	Id string `json:"id"`

	// **参数解释：** 操作类型。 **取值范围：** - UPDATE：修改HRA策略规则。 - DELETE：删除HRA策略规则。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	Operate string `json:"operate"`

	// **参数解释：** 扩缩容类型。 **取值范围：** - SIMULATOR_ALGO：模拟器算法扩缩容类型。 **约束限制：** 不涉及。 **默认取值：** 不涉及。
	ScalerType *string `json:"scaler_type,omitempty"`

	// **参数解释：** SLO配置参数信息。 **取值范围：** 不涉及。
	SloInfo []SloInfo `json:"slo_info"`

	// **参数解释：** 指标信息。 **取值范围：** 不涉及。
	Metrics []Metrics `json:"metrics"`

	// **参数解释：** 角色扩缩策略（不会进行实质扩缩，因此该配置值无效）。 **取值范围：** 1~128。
	RoleReplica *[]RoleReplica `json:"role_replica,omitempty"`
}

func (o HraRuleUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HraRuleUpdateRequest struct{}"
	}

	return strings.Join([]string{"HraRuleUpdateRequest", string(data)}, " ")
}
