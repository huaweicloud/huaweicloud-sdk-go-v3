package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberCheckJobResultGroup **参数解释**：配置检查。
type MemberCheckJobResultGroup struct {

	// **参数解释**：检查结果。  **取值范围**：true表示检查通过，false为检查不通过。
	CheckResult *bool `json:"check_result,omitempty"`

	// **参数解释**：分组检查项汇总。
	CheckItems *[]MemberCheckJobResultItem `json:"check_items,omitempty"`

	// **参数解释**：分组检查任务状态。  **取值范围**： - processed：检查任务执行完成。 - processing：检查任务执行中。 - failed：检查任务执行失败。
	CheckStatus *string `json:"check_status,omitempty"`
}

func (o MemberCheckJobResultGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberCheckJobResultGroup struct{}"
	}

	return strings.Join([]string{"MemberCheckJobResultGroup", string(data)}, " ")
}
