package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BindAgentPolicyRequestInfo struct {

	// **参数解释**： 策略组ID **约束限制**： 必填 **取值范围**： 字符长度1-20位 **默认取值**： 不涉及
	GroupId string `json:"group_id"`

	// **参数解释**： 智能体ID列表 **约束限制**： 必填 **取值范围**： 1-2000个策略组ID **默认取值**： 不涉及
	AgentIdList []string `json:"agent_id_list"`
}

func (o BindAgentPolicyRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindAgentPolicyRequestInfo struct{}"
	}

	return strings.Join([]string{"BindAgentPolicyRequestInfo", string(data)}, " ")
}
