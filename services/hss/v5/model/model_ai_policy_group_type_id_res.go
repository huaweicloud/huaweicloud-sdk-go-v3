package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupTypeIdRes **参数解释**： 策略组类型 **取值范围**： - 0：Agent运行时防护-OfficeAce - 1：Agent行为轨迹检测-OfficeAce - 2：Agent角色限定 - 3：Agent运行时防护-AgentArts - 4：Agent行为轨迹检测-AgentArts - 5：Agent运行时防护-OpenClaw - 6：Agent运行时防护-其他三方Claw
type AiPolicyGroupTypeIdRes struct {
}

func (o AiPolicyGroupTypeIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupTypeIdRes struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupTypeIdRes", string(data)}, " ")
}
