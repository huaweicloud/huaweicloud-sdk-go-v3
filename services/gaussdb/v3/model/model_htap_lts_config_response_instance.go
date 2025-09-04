package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapLtsConfigResponseInstance **参数解释**： HTAP标准版实例信息。  **约束限制**： 不涉及。
type HtapLtsConfigResponseInstance struct {

	// **参数解释**： HTAP标准版实例ID，此参数是实例的唯一标识。  **取值范围**：  只能由英文字母、数字组成，后缀为in17，长度为36个字符。
	Id string `json:"id"`

	// **参数解释**： HTAP标准版实例名称。  **取值范围**：  不涉及。
	Name string `json:"name"`

	// **参数解释**： 引擎类型。  **取值范围**：  不涉及。
	EngineName string `json:"engine_name"`

	// **参数解释**： 引擎版本。  **取值范围**：  不涉及。
	EngineVersion string `json:"engine_version"`

	// **参数解释**： HTAP标准版实例状态。  **取值范围**：  不涉及。
	Status string `json:"status"`

	// **参数解释**： 企业project ID。  **取值范围**：  不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 企业project名称。  **取值范围**：  不涉及。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`
}

func (o HtapLtsConfigResponseInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapLtsConfigResponseInstance struct{}"
	}

	return strings.Join([]string{"HtapLtsConfigResponseInstance", string(data)}, " ")
}
