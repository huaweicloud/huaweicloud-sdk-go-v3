package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiAgentPolicy 查询AI实例绑定策略
type AiAgentPolicy struct {

	// **参数解释**: 策略组名称 **取值范围**: 字符长度1-128位
	PolicyGroupName *string `json:"policy_group_name,omitempty"`

	// **参数解释**： 策略组ID **取值范围**： 字符长度1-20位
	PolicyGroupId *string `json:"policy_group_id,omitempty"`

	// **参数解释**： 策略组类型 **取值范围**： - 0：Agent运行时防护-OfficeAce - 1：Agent行为轨迹检测-OfficeAce - 2：Agent角色限定 - 3：Agent运行时防护-AgentArts - 4：Agent行为轨迹检测-AgentArts - 5：Agent运行时防护-OpenClaw - 6：Agent运行时防护-其他三方Claw
	PolicyGroupType *int32 `json:"policy_group_type,omitempty"`

	// **参数解释**: 策略组是否启用 **取值范围**: - true：已启用 - false：未启用
	PolicyGroupEnabled *bool `json:"policy_group_enabled,omitempty"`

	// **参数解释**: 是否是默认策略 **取值范围**: - false：否 - true：是
	IsDefault *bool `json:"is_default,omitempty"`
}

func (o AiAgentPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAgentPolicy struct{}"
	}

	return strings.Join([]string{"AiAgentPolicy", string(data)}, " ")
}
