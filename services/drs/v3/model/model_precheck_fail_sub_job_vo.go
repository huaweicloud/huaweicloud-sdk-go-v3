package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PrecheckFailSubJobVo 预检查失败子任务信息体
type PrecheckFailSubJobVo struct {

	// **参数解释**：  预检查失败子任务ID，用于唯一标识一个预检查失败的子任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：  预检查失败子任务的名称，用于标识失败的子任务。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：  预检查失败子任务的检查结果，表示该子任务的检查是否通过。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及
	CheckResult *string `json:"check_result,omitempty"`
}

func (o PrecheckFailSubJobVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrecheckFailSubJobVo struct{}"
	}

	return strings.Join([]string{"PrecheckFailSubJobVo", string(data)}, " ")
}
