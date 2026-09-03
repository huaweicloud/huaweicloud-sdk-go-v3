package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProxyEngineVersionInfo 数据库代理节点引擎版本信息
type ProxyEngineVersionInfo struct {

	// **参数解释**：  当前引擎版本。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	CurrentEngineVersion *string `json:"current_engine_version,omitempty"`

	// **参数解释**：  目标引擎版本。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	TargetEngineVersion *string `json:"target_engine_version,omitempty"`

	// **参数解释**：  是否可升级标志。true表示可以升级，false表示不可升级。  **约束限制**：  不涉及。  **取值范围**：  - true - false  **默认取值**：  不涉及。
	UpgradeFlag *bool `json:"upgrade_flag,omitempty"`

	// **参数解释**：  代理节点ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ProxyId *string `json:"proxy_id,omitempty"`

	// **参数解释**：  升级风险列表。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Risks *[]EngineRiskDesc `json:"risks,omitempty"`
}

func (o ProxyEngineVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyEngineVersionInfo struct{}"
	}

	return strings.Join([]string{"ProxyEngineVersionInfo", string(data)}, " ")
}
