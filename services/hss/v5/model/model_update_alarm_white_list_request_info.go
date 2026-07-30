package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlarmWhiteListRequestInfo 编辑告警白名单生效主机
type UpdateAlarmWhiteListRequestInfo struct {

	// **参数解释**： 规则ID **约束限制**： 必填 **取值范围**： 字符长度1-36位 **默认取值**： 不涉及
	RuleId string `json:"rule_id"`

	// **参数解释**: 是否选择所有主机 **约束限制**: 不涉及 **取值范围**: - true：是 - false：否 **默认取值**: false
	Scope *bool `json:"scope,omitempty"`

	// **参数解释**: agent列表 **约束限制**: 不涉及 **取值范围**: 1-1000个agentID **默认取值**: 不涉及
	AgentIds *[]string `json:"agent_ids,omitempty"`

	// **参数解释**: 实例ID列表 **约束限制**: 当需要为serverless配置规则时，传入此字段 **取值范围**: 1-1000个实例ID **默认取值**: 不涉及
	InstanceIds *[]string `json:"instance_ids,omitempty"`
}

func (o UpdateAlarmWhiteListRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmWhiteListRequestInfo struct{}"
	}

	return strings.Join([]string{"UpdateAlarmWhiteListRequestInfo", string(data)}, " ")
}
