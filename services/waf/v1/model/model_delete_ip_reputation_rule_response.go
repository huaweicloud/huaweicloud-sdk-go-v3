package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpReputationRuleResponse Response Object
type DeleteIpReputationRuleResponse struct {

	// **参数解释：** 被删除规则的ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 所属防护策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policyid *string `json:"policyid,omitempty"`

	// **参数解释：** 被删除规则的名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 规则删除时间戳（毫秒级） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释：** 被删除规则的描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则删除时的状态（0表示已关闭/删除） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 规则类型（如idc表示机房IP情报类型） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释：** 规则关联的标签列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags *[]string `json:"tags,omitempty"`

	Action         *DeleteIpReputationRuleResponseBodyAction `json:"action,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o DeleteIpReputationRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpReputationRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpReputationRuleResponse", string(data)}, " ")
}
