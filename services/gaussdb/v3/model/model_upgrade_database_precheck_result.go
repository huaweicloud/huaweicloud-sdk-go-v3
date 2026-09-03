package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeDatabasePrecheckResult **参数解释**：  升级预检查结果。  **取值范围**：  不涉及。
type UpgradeDatabasePrecheckResult struct {

	// **参数解释**：  升级预检查项目。  **取值范围**：  - Upgrade permission check：升级权限检查。 - Instance version check：实例源版本检查。 - Resource check：资源检查。 - Upgrade feature compatibility check：升级特性兼容性检查。
	CheckItem string `json:"check_item"`

	// **参数解释**：  升级预检查项说明。  **取值范围**：  不涉及。
	CheckDescription string `json:"check_description"`

	// **参数解释**：  升级预检查对象。  **取值范围**：  不涉及。
	CheckObject string `json:"check_object"`

	// **参数解释**：  升级预检查项的检查状态。  **取值范围**：  - passed：检查通过。 - failed：检查失败。
	CheckStatus string `json:"check_status"`
}

func (o UpgradeDatabasePrecheckResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeDatabasePrecheckResult struct{}"
	}

	return strings.Join([]string{"UpgradeDatabasePrecheckResult", string(data)}, " ")
}
