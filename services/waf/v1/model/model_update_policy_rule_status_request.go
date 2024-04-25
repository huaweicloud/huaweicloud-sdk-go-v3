package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePolicyRuleStatusRequest Request Object
type UpdatePolicyRuleStatusRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 策略id（策略id从查询防护策略列表接口获取）
	PolicyId string `json:"policy_id"`

	// 策略类型
	Ruletype UpdatePolicyRuleStatusRequestRuletype `json:"ruletype"`

	// 规则id，根据不同的规则类型（ruletype）调用规则列表接口获取规则id，例如黑白名单（whiteblackip）规则id，您可以通过调用查询黑白名单规则列表（ListWhiteblackipRule）获取规则id
	RuleId string `json:"rule_id"`

	Body *UpdatePolicyRuleStatusRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusRequest", string(data)}, " ")
}

type UpdatePolicyRuleStatusRequestRuletype struct {
	value string
}

type UpdatePolicyRuleStatusRequestRuletypeEnum struct {
	WHITEBLACKIP UpdatePolicyRuleStatusRequestRuletype
	GEOIP        UpdatePolicyRuleStatusRequestRuletype
	PRIVACY      UpdatePolicyRuleStatusRequestRuletype
	ANTITAMPER   UpdatePolicyRuleStatusRequestRuletype
	CUSTOM       UpdatePolicyRuleStatusRequestRuletype
	IGNORE       UpdatePolicyRuleStatusRequestRuletype
	CC           UpdatePolicyRuleStatusRequestRuletype
}

func GetUpdatePolicyRuleStatusRequestRuletypeEnum() UpdatePolicyRuleStatusRequestRuletypeEnum {
	return UpdatePolicyRuleStatusRequestRuletypeEnum{
		WHITEBLACKIP: UpdatePolicyRuleStatusRequestRuletype{
			value: "whiteblackip",
		},
		GEOIP: UpdatePolicyRuleStatusRequestRuletype{
			value: "geoip",
		},
		PRIVACY: UpdatePolicyRuleStatusRequestRuletype{
			value: "privacy",
		},
		ANTITAMPER: UpdatePolicyRuleStatusRequestRuletype{
			value: "antitamper",
		},
		CUSTOM: UpdatePolicyRuleStatusRequestRuletype{
			value: "custom",
		},
		IGNORE: UpdatePolicyRuleStatusRequestRuletype{
			value: "ignore",
		},
		CC: UpdatePolicyRuleStatusRequestRuletype{
			value: "cc",
		},
	}
}

func (c UpdatePolicyRuleStatusRequestRuletype) Value() string {
	return c.value
}

func (c UpdatePolicyRuleStatusRequestRuletype) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRuleStatusRequestRuletype) UnmarshalJSON(b []byte) error {
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
