package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuditRuleRisksNewRequest Request Object
type ListAuditRuleRisksNewRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// 风险名称
	Name *string `json:"name,omitempty"`

	// **参数解释**： 风险级别 **约束限制**： 以取值范围为准 **取值范围**： - LOW：低风险 - MEDIUM：中风险 - HIGH：高风险 - NO_RISK：无风险 **默认取值**： 不涉及
	RiskLevels *string `json:"risk_levels,omitempty"`

	// **参数解释**： 实例前端是否支持按数据库类型展示风险规则 **约束限制**： 以取值范围为准 **取值范围**： - true: 支持 - false: 不支持 **默认取值**： false: 不支持
	SupportDbClassifyRule *bool `json:"support_db_classify_rule,omitempty"`
}

func (o ListAuditRuleRisksNewRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleRisksNewRequest struct{}"
	}

	return strings.Join([]string{"ListAuditRuleRisksNewRequest", string(data)}, " ")
}
