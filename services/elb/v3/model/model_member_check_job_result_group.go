package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MemberCheckJobResultGroup **参数解释**：配置检查。
type MemberCheckJobResultGroup struct {

	// **参数解释**：检查结果，true表示检查通过，false为检查不通过。  **取值范围**：不涉及
	CheckResult *bool `json:"check_result,omitempty"`

	CheckItems *[]MemberCheckJobResultItem `json:"check_items,omitempty"`

	// **参数解释**：processed检查完成，processing检查中，failed检查失败。  **取值范围**：不涉及
	CheckStatus *string `json:"check_status,omitempty"`
}

func (o MemberCheckJobResultGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MemberCheckJobResultGroup struct{}"
	}

	return strings.Join([]string{"MemberCheckJobResultGroup", string(data)}, " ")
}
