package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScheduledTaskRequestBody 创建定时任务请求体
type ScheduledTaskRequestBody struct {

	// 选择的四号提权单信息
	TicketInfos *[]TicketInfo `json:"ticket_infos,omitempty"`

	// 任务名称
	Name string `json:"name"`

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 委托名称
	AgencyName *string `json:"agency_name,omitempty"`

	// 版本号
	VersionNo string `json:"version_no"`

	TriggerTime *TriggerTime `json:"trigger_time"`

	// 任务类型
	TaskType ScheduledTaskRequestBodyTaskType `json:"task_type"`

	// 关联任务ID（脚本ID/作业ID）
	AssociatedTaskId string `json:"associated_task_id"`

	// 关联任务的类型
	AssociatedTaskType ScheduledTaskRequestBodyAssociatedTaskType `json:"associated_task_type"`

	// 关联任务名称（脚本名称/作业名称）
	AssociatedTaskName string `json:"associated_task_name"`

	// 关联任务名称(英文)（脚本名称/作业名称）
	AssociatedTaskNameEn *string `json:"associated_task_name_en,omitempty"`

	// 关联任务的企业项目ID
	AssociatedTaskEnterpriseProjectId *string `json:"associated_task_enterprise_project_id,omitempty"`

	// 作业实例模式
	RunbookInstanceMode *ScheduledTaskRequestBodyRunbookInstanceMode `json:"runbook_instance_mode,omitempty"`

	// 风险等级
	RiskLevel ScheduledTaskRequestBodyRiskLevel `json:"risk_level"`

	// 执行参数，值为json串
	InputParam map[string]string `json:"input_param"`

	// 目标实例信息
	TargetInstances []ScheduleInstance `json:"target_instances"`

	// 是否开启入库人工审核
	EnableApprove bool `json:"enable_approve"`

	ReviewerNotification *MessageNotification `json:"reviewer_notification,omitempty"`

	// 审核人昵称
	ReviewerUserName *string `json:"reviewer_user_name,omitempty"`

	// 是否启用消息通知
	EnableMessageNotification bool `json:"enable_message_notification"`

	MessageNotification *MessageNotification `json:"message_notification,omitempty"`
}

func (o ScheduledTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledTaskRequestBody struct{}"
	}

	return strings.Join([]string{"ScheduledTaskRequestBody", string(data)}, " ")
}

type ScheduledTaskRequestBodyTaskType struct {
	value string
}

type ScheduledTaskRequestBodyTaskTypeEnum struct {
	RUNBOOK ScheduledTaskRequestBodyTaskType
	SCRIPT  ScheduledTaskRequestBodyTaskType
}

func GetScheduledTaskRequestBodyTaskTypeEnum() ScheduledTaskRequestBodyTaskTypeEnum {
	return ScheduledTaskRequestBodyTaskTypeEnum{
		RUNBOOK: ScheduledTaskRequestBodyTaskType{
			value: "RUNBOOK",
		},
		SCRIPT: ScheduledTaskRequestBodyTaskType{
			value: "SCRIPT",
		},
	}
}

func (c ScheduledTaskRequestBodyTaskType) Value() string {
	return c.value
}

func (c ScheduledTaskRequestBodyTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledTaskRequestBodyTaskType) UnmarshalJSON(b []byte) error {
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

type ScheduledTaskRequestBodyAssociatedTaskType struct {
	value string
}

type ScheduledTaskRequestBodyAssociatedTaskTypeEnum struct {
	COMMUNAL      ScheduledTaskRequestBodyAssociatedTaskType
	CUSTOMIZATION ScheduledTaskRequestBodyAssociatedTaskType
}

func GetScheduledTaskRequestBodyAssociatedTaskTypeEnum() ScheduledTaskRequestBodyAssociatedTaskTypeEnum {
	return ScheduledTaskRequestBodyAssociatedTaskTypeEnum{
		COMMUNAL: ScheduledTaskRequestBodyAssociatedTaskType{
			value: "COMMUNAL",
		},
		CUSTOMIZATION: ScheduledTaskRequestBodyAssociatedTaskType{
			value: "CUSTOMIZATION",
		},
	}
}

func (c ScheduledTaskRequestBodyAssociatedTaskType) Value() string {
	return c.value
}

func (c ScheduledTaskRequestBodyAssociatedTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledTaskRequestBodyAssociatedTaskType) UnmarshalJSON(b []byte) error {
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

type ScheduledTaskRequestBodyRunbookInstanceMode struct {
	value string
}

type ScheduledTaskRequestBodyRunbookInstanceModeEnum struct {
	SAME ScheduledTaskRequestBodyRunbookInstanceMode
	DIFF ScheduledTaskRequestBodyRunbookInstanceMode
}

func GetScheduledTaskRequestBodyRunbookInstanceModeEnum() ScheduledTaskRequestBodyRunbookInstanceModeEnum {
	return ScheduledTaskRequestBodyRunbookInstanceModeEnum{
		SAME: ScheduledTaskRequestBodyRunbookInstanceMode{
			value: "SAME",
		},
		DIFF: ScheduledTaskRequestBodyRunbookInstanceMode{
			value: "DIFF",
		},
	}
}

func (c ScheduledTaskRequestBodyRunbookInstanceMode) Value() string {
	return c.value
}

func (c ScheduledTaskRequestBodyRunbookInstanceMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledTaskRequestBodyRunbookInstanceMode) UnmarshalJSON(b []byte) error {
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

type ScheduledTaskRequestBodyRiskLevel struct {
	value string
}

type ScheduledTaskRequestBodyRiskLevelEnum struct {
	HIGH   ScheduledTaskRequestBodyRiskLevel
	MEDIUM ScheduledTaskRequestBodyRiskLevel
	LOW    ScheduledTaskRequestBodyRiskLevel
}

func GetScheduledTaskRequestBodyRiskLevelEnum() ScheduledTaskRequestBodyRiskLevelEnum {
	return ScheduledTaskRequestBodyRiskLevelEnum{
		HIGH: ScheduledTaskRequestBodyRiskLevel{
			value: "HIGH",
		},
		MEDIUM: ScheduledTaskRequestBodyRiskLevel{
			value: "MEDIUM",
		},
		LOW: ScheduledTaskRequestBodyRiskLevel{
			value: "LOW",
		},
	}
}

func (c ScheduledTaskRequestBodyRiskLevel) Value() string {
	return c.value
}

func (c ScheduledTaskRequestBodyRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledTaskRequestBodyRiskLevel) UnmarshalJSON(b []byte) error {
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
