package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScriptPropertiesModel 脚本标签：风险等级risk_level(取值范围：LOW/MEDIUM/HIGH)、审批人reviewer、脚本解释器interpreter等
type ScriptPropertiesModel struct {

	// 风险等级 LOW:低风险 MEDIUM:中风险 HIGH:高风险
	RiskLevel ScriptPropertiesModelRiskLevel `json:"risk_level"`

	// 脚本版本号
	Version string `json:"version"`

	// 审批人，不填写不需要审批
	Reviewers *[]ReviewerInfo `json:"reviewers,omitempty"`

	// 审批消息通知协议，用于通知审批人 DEFAULT：默认 SMS：短信 EMAIL：邮件 DING_TALK：钉钉 WE_LINK：welink WECHAT:微信 CALLNOTIFY：语言 NOT_TO_NOTIFY：不通知
	Protocol *ScriptPropertiesModelProtocol `json:"protocol,omitempty"`
}

func (o ScriptPropertiesModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScriptPropertiesModel struct{}"
	}

	return strings.Join([]string{"ScriptPropertiesModel", string(data)}, " ")
}

type ScriptPropertiesModelRiskLevel struct {
	value string
}

type ScriptPropertiesModelRiskLevelEnum struct {
	LOW    ScriptPropertiesModelRiskLevel
	MEDIUM ScriptPropertiesModelRiskLevel
	HIGH   ScriptPropertiesModelRiskLevel
}

func GetScriptPropertiesModelRiskLevelEnum() ScriptPropertiesModelRiskLevelEnum {
	return ScriptPropertiesModelRiskLevelEnum{
		LOW: ScriptPropertiesModelRiskLevel{
			value: "LOW",
		},
		MEDIUM: ScriptPropertiesModelRiskLevel{
			value: "MEDIUM",
		},
		HIGH: ScriptPropertiesModelRiskLevel{
			value: "HIGH",
		},
	}
}

func (c ScriptPropertiesModelRiskLevel) Value() string {
	return c.value
}

func (c ScriptPropertiesModelRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptPropertiesModelRiskLevel) UnmarshalJSON(b []byte) error {
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

type ScriptPropertiesModelProtocol struct {
	value string
}

type ScriptPropertiesModelProtocolEnum struct {
	DEFAULT       ScriptPropertiesModelProtocol
	SMS           ScriptPropertiesModelProtocol
	EMAIL         ScriptPropertiesModelProtocol
	DING_TALK     ScriptPropertiesModelProtocol
	WE_LINK       ScriptPropertiesModelProtocol
	WECHAT        ScriptPropertiesModelProtocol
	CALLNOTIFY    ScriptPropertiesModelProtocol
	NOT_TO_NOTIFY ScriptPropertiesModelProtocol
}

func GetScriptPropertiesModelProtocolEnum() ScriptPropertiesModelProtocolEnum {
	return ScriptPropertiesModelProtocolEnum{
		DEFAULT: ScriptPropertiesModelProtocol{
			value: "DEFAULT",
		},
		SMS: ScriptPropertiesModelProtocol{
			value: "SMS",
		},
		EMAIL: ScriptPropertiesModelProtocol{
			value: "EMAIL",
		},
		DING_TALK: ScriptPropertiesModelProtocol{
			value: "DING_TALK",
		},
		WE_LINK: ScriptPropertiesModelProtocol{
			value: "WE_LINK",
		},
		WECHAT: ScriptPropertiesModelProtocol{
			value: "WECHAT",
		},
		CALLNOTIFY: ScriptPropertiesModelProtocol{
			value: "CALLNOTIFY",
		},
		NOT_TO_NOTIFY: ScriptPropertiesModelProtocol{
			value: "NOT_TO_NOTIFY",
		},
	}
}

func (c ScriptPropertiesModelProtocol) Value() string {
	return c.value
}

func (c ScriptPropertiesModelProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScriptPropertiesModelProtocol) UnmarshalJSON(b []byte) error {
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
