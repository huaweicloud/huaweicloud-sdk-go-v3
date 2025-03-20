package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SlaRecord Sla 规则记录
type SlaRecord struct {

	// 状态类型: EVENT_TICKET_NOT_ACCEPTED EVENT_TICKET_PROCESSING EVENT_TICKET_NOT_IN_TIME EVENT_TICKET_RESOLVED ALARM_TICKET_RESOLVED ALARM_TICKET_ALARMING ALARM_TICKET_NOT_IN_TIME CHANGE_NOT_IN_TIME CHANGE_REVIEW CHANGE_IMPLEMENTATION CHANGE_VERIFICATION TO_DO_TASKS_NOT_IN_TIME TO_DO_TASKS_TO_BE_HANDLED TO_DO_TASKS_PROCESSING TO_DO_TASKS_COMPLETED ISSUE_TICKET_NOT_IN_TIME ISSUE_TICKET_NOT_ACCEPTED ISSUE_TICKET_POSITIONING ISSUE_TICKET_WAITING_IMPLEMENT ISSUE_TICKET_RESOLVED
	StatusType *SlaRecordStatusType `json:"status_type,omitempty"`

	// 子状态(NORMAL,FORWARDING,RESUBMIT)
	SubTriggerNode *SlaRecordSubTriggerNode `json:"sub_trigger_node,omitempty"`

	// Sla状态  false 未打破
	SlaStatus *bool `json:"sla_status,omitempty"`

	// 工单ID
	SlaOrderId *string `json:"sla_order_id,omitempty"`

	// 持续时间
	Duration *int64 `json:"duration,omitempty"`

	// SLA 触发规则是否开启
	TriggeringRuleEnable *bool `json:"triggering_rule_enable,omitempty"`

	TriggeringRule *SlaTriggeringRuleInfo `json:"triggering_rule,omitempty"`

	// 打破时间
	BreakTime *string `json:"break_time,omitempty"`

	// 对象ID
	OwnerId *string `json:"owner_id,omitempty"`

	// 对象人名
	OwnerName *string `json:"owner_name,omitempty"`

	// 通知时间
	NoticeTime *string `json:"notice_time,omitempty"`

	// 状态开始时间
	StatusStartTime *string `json:"status_start_time,omitempty"`

	// SLA 预警通知是否开启
	PreWarningRuleEnable *bool `json:"pre_warning_rule_enable,omitempty"`

	PreWarningRule *SlaPreWarningRuleInfo `json:"pre_warning_rule,omitempty"`

	// SLA 上升通知是否开启
	EscalateRuleEnable *bool `json:"escalate_rule_enable,omitempty"`

	// SLA 上升通知配置
	EscalateRules *[]SlaEscalationRuleInfo `json:"escalate_rules,omitempty"`

	// SLA是否在当前状态
	CurrentNode *bool `json:"current_node,omitempty"`
}

func (o SlaRecord) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlaRecord struct{}"
	}

	return strings.Join([]string{"SlaRecord", string(data)}, " ")
}

type SlaRecordStatusType struct {
	value string
}

type SlaRecordStatusTypeEnum struct {
	EVENT_TICKET_NOT_ACCEPTED      SlaRecordStatusType
	EVENT_TICKET_PROCESSING        SlaRecordStatusType
	EVENT_TICKET_NOT_IN_TIME       SlaRecordStatusType
	EVENT_TICKET_RESOLVED          SlaRecordStatusType
	ALARM_TICKET_RESOLVED          SlaRecordStatusType
	ALARM_TICKET_ALARMING          SlaRecordStatusType
	ALARM_TICKET_NOT_IN_TIME       SlaRecordStatusType
	CHANGE_NOT_IN_TIME             SlaRecordStatusType
	CHANGE_REVIEW                  SlaRecordStatusType
	CHANGE_IMPLEMENTATION          SlaRecordStatusType
	CHANGE_VERIFICATION            SlaRecordStatusType
	TO_DO_TASKS_NOT_IN_TIME        SlaRecordStatusType
	TO_DO_TASKS_TO_BE_HANDLED      SlaRecordStatusType
	TO_DO_TASKS_PROCESSING         SlaRecordStatusType
	TO_DO_TASKS_COMPLETED          SlaRecordStatusType
	ISSUE_TICKET_NOT_IN_TIME       SlaRecordStatusType
	ISSUE_TICKET_NOT_ACCEPTED      SlaRecordStatusType
	ISSUE_TICKET_POSITIONING       SlaRecordStatusType
	ISSUE_TICKET_WAITING_IMPLEMENT SlaRecordStatusType
	ISSUE_TICKET_RESOLVED          SlaRecordStatusType
}

func GetSlaRecordStatusTypeEnum() SlaRecordStatusTypeEnum {
	return SlaRecordStatusTypeEnum{
		EVENT_TICKET_NOT_ACCEPTED: SlaRecordStatusType{
			value: "EVENT_TICKET_NOT_ACCEPTED",
		},
		EVENT_TICKET_PROCESSING: SlaRecordStatusType{
			value: "EVENT_TICKET_PROCESSING",
		},
		EVENT_TICKET_NOT_IN_TIME: SlaRecordStatusType{
			value: "EVENT_TICKET_NOT_IN_TIME",
		},
		EVENT_TICKET_RESOLVED: SlaRecordStatusType{
			value: "EVENT_TICKET_RESOLVED",
		},
		ALARM_TICKET_RESOLVED: SlaRecordStatusType{
			value: "ALARM_TICKET_RESOLVED",
		},
		ALARM_TICKET_ALARMING: SlaRecordStatusType{
			value: "ALARM_TICKET_ALARMING",
		},
		ALARM_TICKET_NOT_IN_TIME: SlaRecordStatusType{
			value: "ALARM_TICKET_NOT_IN_TIME",
		},
		CHANGE_NOT_IN_TIME: SlaRecordStatusType{
			value: "CHANGE_NOT_IN_TIME",
		},
		CHANGE_REVIEW: SlaRecordStatusType{
			value: "CHANGE_REVIEW",
		},
		CHANGE_IMPLEMENTATION: SlaRecordStatusType{
			value: "CHANGE_IMPLEMENTATION",
		},
		CHANGE_VERIFICATION: SlaRecordStatusType{
			value: "CHANGE_VERIFICATION",
		},
		TO_DO_TASKS_NOT_IN_TIME: SlaRecordStatusType{
			value: "TO_DO_TASKS_NOT_IN_TIME",
		},
		TO_DO_TASKS_TO_BE_HANDLED: SlaRecordStatusType{
			value: "TO_DO_TASKS_TO_BE_HANDLED",
		},
		TO_DO_TASKS_PROCESSING: SlaRecordStatusType{
			value: "TO_DO_TASKS_PROCESSING",
		},
		TO_DO_TASKS_COMPLETED: SlaRecordStatusType{
			value: "TO_DO_TASKS_COMPLETED",
		},
		ISSUE_TICKET_NOT_IN_TIME: SlaRecordStatusType{
			value: "ISSUE_TICKET_NOT_IN_TIME",
		},
		ISSUE_TICKET_NOT_ACCEPTED: SlaRecordStatusType{
			value: "ISSUE_TICKET_NOT_ACCEPTED",
		},
		ISSUE_TICKET_POSITIONING: SlaRecordStatusType{
			value: "ISSUE_TICKET_POSITIONING",
		},
		ISSUE_TICKET_WAITING_IMPLEMENT: SlaRecordStatusType{
			value: "ISSUE_TICKET_WAITING_IMPLEMENT",
		},
		ISSUE_TICKET_RESOLVED: SlaRecordStatusType{
			value: "ISSUE_TICKET_RESOLVED",
		},
	}
}

func (c SlaRecordStatusType) Value() string {
	return c.value
}

func (c SlaRecordStatusType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlaRecordStatusType) UnmarshalJSON(b []byte) error {
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

type SlaRecordSubTriggerNode struct {
	value string
}

type SlaRecordSubTriggerNodeEnum struct {
	NORMAL     SlaRecordSubTriggerNode
	FORWARDING SlaRecordSubTriggerNode
	RESUBMIT   SlaRecordSubTriggerNode
}

func GetSlaRecordSubTriggerNodeEnum() SlaRecordSubTriggerNodeEnum {
	return SlaRecordSubTriggerNodeEnum{
		NORMAL: SlaRecordSubTriggerNode{
			value: "NORMAL",
		},
		FORWARDING: SlaRecordSubTriggerNode{
			value: "FORWARDING",
		},
		RESUBMIT: SlaRecordSubTriggerNode{
			value: "RESUBMIT",
		},
	}
}

func (c SlaRecordSubTriggerNode) Value() string {
	return c.value
}

func (c SlaRecordSubTriggerNode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SlaRecordSubTriggerNode) UnmarshalJSON(b []byte) error {
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
