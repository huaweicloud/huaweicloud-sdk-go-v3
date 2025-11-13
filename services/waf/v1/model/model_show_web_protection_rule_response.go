package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWebProtectionRuleResponse Response Object
type ShowWebProtectionRuleResponse struct {

	// **参数解释：** 规则id，唯一标识一条Web防护规则 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释：** 关联的CVE编号，对应规则防护的漏洞的CVE标准编号（如CVE-2021-41773） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CveNumber *string `json:"cve_number,omitempty"`

	// **参数解释：** 危险等级 - 1：高危 - 2：中危 - 3：低危  **取值范围：** - 1 - 2 - 3
	RiskLevel *ShowWebProtectionRuleResponseRiskLevel `json:"risk_level,omitempty"`

	// **参数解释：** 应用类型，指定规则适用的应用场景（如Web应用、API接口） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ApplicationType *string `json:"application_type,omitempty"`

	// **参数解释：** 防护类型，区分规则的防护类别（如SQL注入防护、XSS防护、命令注入防护） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	ProtectionType *string `json:"protection_type,omitempty"`

	// **参数解释：** 描述，对Web防护规则的功能、适用场景等补充说明 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释：** 规则上线时间，Web防护规则正式启用的时间（时间戳格式） **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	CreateTime     *int64 `json:"create_time,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowWebProtectionRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWebProtectionRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowWebProtectionRuleResponse", string(data)}, " ")
}

type ShowWebProtectionRuleResponseRiskLevel struct {
	value int32
}

type ShowWebProtectionRuleResponseRiskLevelEnum struct {
	E_1 ShowWebProtectionRuleResponseRiskLevel
	E_2 ShowWebProtectionRuleResponseRiskLevel
	E_3 ShowWebProtectionRuleResponseRiskLevel
}

func GetShowWebProtectionRuleResponseRiskLevelEnum() ShowWebProtectionRuleResponseRiskLevelEnum {
	return ShowWebProtectionRuleResponseRiskLevelEnum{
		E_1: ShowWebProtectionRuleResponseRiskLevel{
			value: 1,
		}, E_2: ShowWebProtectionRuleResponseRiskLevel{
			value: 2,
		}, E_3: ShowWebProtectionRuleResponseRiskLevel{
			value: 3,
		},
	}
}

func (c ShowWebProtectionRuleResponseRiskLevel) Value() int32 {
	return c.value
}

func (c ShowWebProtectionRuleResponseRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWebProtectionRuleResponseRiskLevel) UnmarshalJSON(b []byte) error {
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
