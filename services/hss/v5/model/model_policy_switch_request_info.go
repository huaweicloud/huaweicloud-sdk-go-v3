package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicySwitchRequestInfo struct {

	// **参数解释**： 策略名称 **约束限制**： 必填 **取值范围**： - sp_feature：自保护策略。  **默认取值**： 不涉及
	PolicyName string `json:"policy_name"`

	// **参数解释**: 策略总开关是否开启 **约束限制**: 不涉及 **取值范围**: - true：是。 - false：否。  **默认取值**: 不涉及
	Enable bool `json:"enable"`
}

func (o PolicySwitchRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicySwitchRequestInfo struct{}"
	}

	return strings.Join([]string{"PolicySwitchRequestInfo", string(data)}, " ")
}
