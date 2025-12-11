package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyEngineRisk struct {

	// **参数解释**：  风险级别。  **取值范围**： 1：一级风险。
	Level *int32 `json:"level,omitempty"`

	// **参数解释**：  建议升级原因。  **取值范围**：  不涉及。
	Suggest *string `json:"suggest,omitempty"`

	// **参数解释**：  升级影响。  **取值范围**：  不涉及。
	Influence *string `json:"influence,omitempty"`

	// **参数解释**：  升级指导。  **取值范围**：  不涉及。
	Guidance *string `json:"guidance,omitempty"`

	// **参数解释**：  预估业务影响时长。  **取值范围**：  不涉及。
	ServiceImpactDuration *string `json:"service_impact_duration,omitempty"`

	// **参数解释**：  预估升级时长。  **取值范围**：  不涉及。
	UpgradeDuration *string `json:"upgrade_duration,omitempty"`
}

func (o ProxyEngineRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyEngineRisk struct{}"
	}

	return strings.Join([]string{"ProxyEngineRisk", string(data)}, " ")
}
