package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SlaRuleInfo Sla 的规则
type SlaRuleInfo struct {

	// 状态类型: EVENT_TICKET_NOT_ACCEPTED EVENT_TICKET_PROCESSING EVENT_TICKET_RESOLVED ALARM_TICKET_ALARMING CHANGE_REVIEW CHANGE_IMPLEMENTATION CHANGE_VERIFICATION TO_DO_TASKS_TO_BE_HANDLED TO_DO_TASKS_PROCESSING ISSUE_TICKET_NOT_ACCEPTED ISSUE_TICKET_POSITIONING ISSUE_TICKET_WAITING_IMPLEMENT ISSUE_TICKET_RESOLVED
	StatusType SlaRuleInfoStatusType `json:"status_type"`

	// 触发规则是否生效
	TriggeringRuleEnable bool `json:"triggering_rule_enable"`

	// 预警规则是否生效
	PreWarningRuleEnable bool `json:"pre_warning_rule_enable"`

	// 上升规则是否生效
	EscalateRuleEnable *bool `json:"escalate_rule_enable,omitempty"`

	TriggeringRule *SlaTriggeringRuleInfo `json:"triggering_rule,omitempty"`

	PreWarningRule *SlaPreWarningRuleInfo `json:"pre_warning_rule,omitempty"`

	// 通知配置
	EscalateRules *[]SlaEscalationRuleInfo `json:"escalate_rules,omitempty"`
}

func (o SlaRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaRuleInfo struct{}"
	}

	return strings.Join([]string{"SlaRuleInfo", string(data)}, " ")
}

type SlaRuleInfoStatusType struct {
	value string
}

type SlaRuleInfoStatusTypeEnum struct {
	EVENT_TICKET_NOT_ACCEPTED      SlaRuleInfoStatusType
	EVENT_TICKET_PROCESSING        SlaRuleInfoStatusType
	EVENT_TICKET_RESOLVED          SlaRuleInfoStatusType
	ALARM_TICKET_ALARMING          SlaRuleInfoStatusType
	CHANGE_REVIEW                  SlaRuleInfoStatusType
	CHANGE_IMPLEMENTATION          SlaRuleInfoStatusType
	CHANGE_VERIFICATION            SlaRuleInfoStatusType
	TO_DO_TASKS_TO_BE_HANDLED      SlaRuleInfoStatusType
	TO_DO_TASKS_PROCESSING         SlaRuleInfoStatusType
	ISSUE_TICKET_NOT_ACCEPTED      SlaRuleInfoStatusType
	ISSUE_TICKET_POSITIONING       SlaRuleInfoStatusType
	ISSUE_TICKET_WAITING_IMPLEMENT SlaRuleInfoStatusType
	ISSUE_TICKET_RESOLVED          SlaRuleInfoStatusType
}

func GetSlaRuleInfoStatusTypeEnum() SlaRuleInfoStatusTypeEnum {
	return SlaRuleInfoStatusTypeEnum{
		EVENT_TICKET_NOT_ACCEPTED: SlaRuleInfoStatusType{
			value: "EVENT_TICKET_NOT_ACCEPTED",
		},
		EVENT_TICKET_PROCESSING: SlaRuleInfoStatusType{
			value: "EVENT_TICKET_PROCESSING",
		},
		EVENT_TICKET_RESOLVED: SlaRuleInfoStatusType{
			value: "EVENT_TICKET_RESOLVED",
		},
		ALARM_TICKET_ALARMING: SlaRuleInfoStatusType{
			value: "ALARM_TICKET_ALARMING",
		},
		CHANGE_REVIEW: SlaRuleInfoStatusType{
			value: "CHANGE_REVIEW",
		},
		CHANGE_IMPLEMENTATION: SlaRuleInfoStatusType{
			value: "CHANGE_IMPLEMENTATION",
		},
		CHANGE_VERIFICATION: SlaRuleInfoStatusType{
			value: "CHANGE_VERIFICATION",
		},
		TO_DO_TASKS_TO_BE_HANDLED: SlaRuleInfoStatusType{
			value: "TO_DO_TASKS_TO_BE_HANDLED",
		},
		TO_DO_TASKS_PROCESSING: SlaRuleInfoStatusType{
			value: "TO_DO_TASKS_PROCESSING",
		},
		ISSUE_TICKET_NOT_ACCEPTED: SlaRuleInfoStatusType{
			value: "ISSUE_TICKET_NOT_ACCEPTED",
		},
		ISSUE_TICKET_POSITIONING: SlaRuleInfoStatusType{
			value: "ISSUE_TICKET_POSITIONING",
		},
		ISSUE_TICKET_WAITING_IMPLEMENT: SlaRuleInfoStatusType{
			value: "ISSUE_TICKET_WAITING_IMPLEMENT",
		},
		ISSUE_TICKET_RESOLVED: SlaRuleInfoStatusType{
			value: "ISSUE_TICKET_RESOLVED",
		},
	}
}

func (c SlaRuleInfoStatusType) Value() string {
	return c.value
}

func (c SlaRuleInfoStatusType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlaRuleInfoStatusType) UnmarshalJSON(b []byte) error {
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
