package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertPolicyDto 告警策略DTO
type AlertPolicyDto struct {

	// **参数解释**： 策略ID。 **约束限制**： 更新时传，新增时不传。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PolicyId *string `json:"policyId,omitempty"`

	// **参数解释**： 策略名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 租户ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DomainId *string `json:"domainId,omitempty"`

	// **参数解释**： 项目ID。 **约束限制**： 暂时不用。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId *string `json:"projectId,omitempty"`

	// **参数解释**： 是否为默认策略。 **约束限制**： 不涉及。 **取值范围**： - true：是默认策略。 - false：不是默认策略。 **默认取值**： 不涉及。
	IsDefault *bool `json:"isDefault,omitempty"`

	// **参数解释**： 创建时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	CreateTime *int64 `json:"createTime,omitempty"`

	// **参数解释**： 所有规则。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Rules *[]AlertRuleDto `json:"rules,omitempty"`
}

func (o AlertPolicyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertPolicyDto struct{}"
	}

	return strings.Join([]string{"AlertPolicyDto", string(data)}, " ")
}
