package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WebBasicProtectionRulesItem web基础防护内置规则
type WebBasicProtectionRulesItem struct {

	// **参数解释：** 规则的ID，规则的唯一标识 **取值范围：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 规则描述 **取值范围：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** CVE编号 **取值范围：** 不涉及
	CveNumber *string `json:"cve_number,omitempty"`

	// **参数解释：** 危险等级 - 1：高危 - 2：中危 - 3：低危  **取值范围：** - 1 - 2 - 3
	RiskLevel *WebBasicProtectionRulesItemRiskLevel `json:"risk_level,omitempty"`

	// **参数解释：** 应用类型 **取值范围：** 不涉及
	ApplicationType *string `json:"application_type,omitempty"`

	// **参数解释：** 防护类型 **取值范围：** 不涉及
	ProtectionType *string `json:"protection_type,omitempty"`

	// **参数解释：** 生效时间 **取值范围：** 不涉及
	EffectiveTime *int64 `json:"effective_time,omitempty"`

	// **参数解释：** 创建时间 **取值范围：** 不涉及
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释：** 更新时间 **取值范围：** 不涉及
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o WebBasicProtectionRulesItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebBasicProtectionRulesItem struct{}"
	}

	return strings.Join([]string{"WebBasicProtectionRulesItem", string(data)}, " ")
}

type WebBasicProtectionRulesItemRiskLevel struct {
	value int32
}

type WebBasicProtectionRulesItemRiskLevelEnum struct {
	E_1 WebBasicProtectionRulesItemRiskLevel
	E_2 WebBasicProtectionRulesItemRiskLevel
	E_3 WebBasicProtectionRulesItemRiskLevel
}

func GetWebBasicProtectionRulesItemRiskLevelEnum() WebBasicProtectionRulesItemRiskLevelEnum {
	return WebBasicProtectionRulesItemRiskLevelEnum{
		E_1: WebBasicProtectionRulesItemRiskLevel{
			value: 1,
		}, E_2: WebBasicProtectionRulesItemRiskLevel{
			value: 2,
		}, E_3: WebBasicProtectionRulesItemRiskLevel{
			value: 3,
		},
	}
}

func (c WebBasicProtectionRulesItemRiskLevel) Value() int32 {
	return c.value
}

func (c WebBasicProtectionRulesItemRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WebBasicProtectionRulesItemRiskLevel) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
