package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EngineRiskDesc 引擎风险描述
type EngineRiskDesc struct {

	// **参数解释**：  代理节点ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  引擎名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EngineName *string `json:"engine_name,omitempty"`

	// **参数解释**：  引擎版本。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**：  风险等级（该字段当前无效，默认为1）。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Level *int32 `json:"level,omitempty"`

	// **参数解释**：  建议。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Suggest *string `json:"suggest,omitempty"`

	// **参数解释**：  影响。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Influence *string `json:"influence,omitempty"`

	// **参数解释**：  指导。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Guidance *string `json:"guidance,omitempty"`

	// **参数解释**：  服务影响时长说明。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ServiceImpactDuration *string `json:"service_impact_duration,omitempty"`

	// **参数解释**：  升级时长说明。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	UpgradeDuration *string `json:"upgrade_duration,omitempty"`
}

func (o EngineRiskDesc) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EngineRiskDesc struct{}"
	}

	return strings.Join([]string{"EngineRiskDesc", string(data)}, " ")
}
