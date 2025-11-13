package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConfirmIpReputationRuleResponse Response Object
type ConfirmIpReputationRuleResponse struct {

	// **参数解释：** 规则ID，唯一标识该规则 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 所属防护策略ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Policyid *string `json:"policyid,omitempty"`

	// **参数解释：** 规则名称 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释：** 规则最后更新时间戳（毫秒级） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释：** 规则描述 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则状态（1表示开启，0表示关闭） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Status *int32 `json:"status,omitempty"`

	// **参数解释：** 规则类型（如idc表示机房IP情报类型） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释：** 规则关联的标签列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Tags *[]string `json:"tags,omitempty"`

	Action         *CreateIpReputationRuleResponseBodyAction `json:"action,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ConfirmIpReputationRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmIpReputationRuleResponse struct{}"
	}

	return strings.Join([]string{"ConfirmIpReputationRuleResponse", string(data)}, " ")
}
