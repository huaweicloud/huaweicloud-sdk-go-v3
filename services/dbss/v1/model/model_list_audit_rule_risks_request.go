package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAuditRuleRisksRequest Request Object
type ListAuditRuleRisksRequest struct {

	// **参数解释**： 实例ID。可通过查询实例列表接口ID字段获取 **约束限制**： 不涉及 **取值范围**： 以查询实例列表接口值为准，字符长度32-64。 **默认取值**： 不涉及
	InstanceId string `json:"instance_id"`

	// 风险名称
	Name *string `json:"name,omitempty"`

	// 风险级别 - LOW - MEDIUM - HIGH - NO_RISK
	RiskLevels *ListAuditRuleRisksRequestRiskLevels `json:"risk_levels,omitempty"`
}

func (o ListAuditRuleRisksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuditRuleRisksRequest struct{}"
	}

	return strings.Join([]string{"ListAuditRuleRisksRequest", string(data)}, " ")
}

type ListAuditRuleRisksRequestRiskLevels struct {
	value string
}

type ListAuditRuleRisksRequestRiskLevelsEnum struct {
	LOW     ListAuditRuleRisksRequestRiskLevels
	MEDIUM  ListAuditRuleRisksRequestRiskLevels
	HIGH    ListAuditRuleRisksRequestRiskLevels
	NO_RISK ListAuditRuleRisksRequestRiskLevels
}

func GetListAuditRuleRisksRequestRiskLevelsEnum() ListAuditRuleRisksRequestRiskLevelsEnum {
	return ListAuditRuleRisksRequestRiskLevelsEnum{
		LOW: ListAuditRuleRisksRequestRiskLevels{
			value: "LOW",
		},
		MEDIUM: ListAuditRuleRisksRequestRiskLevels{
			value: "MEDIUM",
		},
		HIGH: ListAuditRuleRisksRequestRiskLevels{
			value: "HIGH",
		},
		NO_RISK: ListAuditRuleRisksRequestRiskLevels{
			value: "NO_RISK",
		},
	}
}

func (c ListAuditRuleRisksRequestRiskLevels) Value() string {
	return c.value
}

func (c ListAuditRuleRisksRequestRiskLevels) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAuditRuleRisksRequestRiskLevels) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
