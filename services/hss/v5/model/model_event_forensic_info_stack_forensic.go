package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EventForensicInfoStackForensic **参数解释**： 堆栈取证信息 **取值范围**： 不涉及
type EventForensicInfoStackForensic struct {

	// **参数解释**： 攻击载荷 **取值范围**： 不涉及
	AttackInputValue *string `json:"attack_input_value,omitempty"`

	// **参数解释**： 堆栈信息 **取值范围**： 不涉及
	AppStack *string `json:"app_stack,omitempty"`

	// **参数解释**： 攻击探针 **取值范围**： 不涉及
	ChkProbe *int32 `json:"chk_probe,omitempty"`

	// **参数解释**： 特性规则 **取值范围**： 不涉及
	ChkRule *int32 `json:"chk_rule,omitempty"`

	// **参数解释**： 检测规则 **取值范围**： 不涉及
	PluginName *int32 `json:"plugin_name,omitempty"`
}

func (o EventForensicInfoStackForensic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventForensicInfoStackForensic struct{}"
	}

	return strings.Join([]string{"EventForensicInfoStackForensic", string(data)}, " ")
}
