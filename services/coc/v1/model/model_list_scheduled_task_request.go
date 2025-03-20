package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListScheduledTaskRequest Request Object
type ListScheduledTaskRequest struct {

	// 企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 定时类型(ONCE,PERIODIC,CRON)
	ScheduledType *ListScheduledTaskRequestScheduledType `json:"scheduled_type,omitempty"`

	// 引用任务类型(SCRIPT,RUNBOOK)
	TaskType *ListScheduledTaskRequestTaskType `json:"task_type,omitempty"`

	// 任务类型(CUSTOMIZATION,COMMUNAL)
	AssociatedTaskType *ListScheduledTaskRequestAssociatedTaskType `json:"associated_task_type,omitempty"`

	// 风险等级
	RiskLevel *string `json:"risk_level,omitempty"`

	// 创建人
	CreatedBy *string `json:"created_by,omitempty"`

	// 审批人ID
	Reviewer *string `json:"reviewer,omitempty"`

	// 审批人昵称
	ReviewerUserName *string `json:"reviewer_user_name,omitempty"`

	// 审批状态(PENDING,REJECTED,PASSED)
	ApproveStatus *ListScheduledTaskRequestApproveStatus `json:"approve_status,omitempty"`

	// 最近执行状态(READY,PROCESSING,ABNORMAL,FINISHED,PAUSED,CANCELED)
	LastExecutionStatus *ListScheduledTaskRequestLastExecutionStatus `json:"last_execution_status,omitempty"`

	// 最近执行时间的查询开始时间
	LastExecutionStartTime *int64 `json:"last_execution_start_time,omitempty"`

	// 最近执行时间的查询结束时间
	LastExecutionEndTime *int64 `json:"last_execution_end_time,omitempty"`

	// 上一页数据的最后一条记录的id
	Marker *string `json:"marker,omitempty"`

	// 区域
	RegionId *string `json:"region_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 分页指针
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"ListScheduledTaskRequest", string(data)}, " ")
}

type ListScheduledTaskRequestScheduledType struct {
	value string
}

type ListScheduledTaskRequestScheduledTypeEnum struct {
	ONCE     ListScheduledTaskRequestScheduledType
	PERIODIC ListScheduledTaskRequestScheduledType
	CRON     ListScheduledTaskRequestScheduledType
}

func GetListScheduledTaskRequestScheduledTypeEnum() ListScheduledTaskRequestScheduledTypeEnum {
	return ListScheduledTaskRequestScheduledTypeEnum{
		ONCE: ListScheduledTaskRequestScheduledType{
			value: "ONCE",
		},
		PERIODIC: ListScheduledTaskRequestScheduledType{
			value: "PERIODIC",
		},
		CRON: ListScheduledTaskRequestScheduledType{
			value: "CRON",
		},
	}
}

func (c ListScheduledTaskRequestScheduledType) Value() string {
	return c.value
}

func (c ListScheduledTaskRequestScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskRequestScheduledType) UnmarshalJSON(b []byte) error {
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

type ListScheduledTaskRequestTaskType struct {
	value string
}

type ListScheduledTaskRequestTaskTypeEnum struct {
	SCRIPT  ListScheduledTaskRequestTaskType
	RUNBOOK ListScheduledTaskRequestTaskType
}

func GetListScheduledTaskRequestTaskTypeEnum() ListScheduledTaskRequestTaskTypeEnum {
	return ListScheduledTaskRequestTaskTypeEnum{
		SCRIPT: ListScheduledTaskRequestTaskType{
			value: "SCRIPT",
		},
		RUNBOOK: ListScheduledTaskRequestTaskType{
			value: "RUNBOOK",
		},
	}
}

func (c ListScheduledTaskRequestTaskType) Value() string {
	return c.value
}

func (c ListScheduledTaskRequestTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskRequestTaskType) UnmarshalJSON(b []byte) error {
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

type ListScheduledTaskRequestAssociatedTaskType struct {
	value string
}

type ListScheduledTaskRequestAssociatedTaskTypeEnum struct {
	CUSTOMIZATION ListScheduledTaskRequestAssociatedTaskType
	COMMUNAL      ListScheduledTaskRequestAssociatedTaskType
}

func GetListScheduledTaskRequestAssociatedTaskTypeEnum() ListScheduledTaskRequestAssociatedTaskTypeEnum {
	return ListScheduledTaskRequestAssociatedTaskTypeEnum{
		CUSTOMIZATION: ListScheduledTaskRequestAssociatedTaskType{
			value: "CUSTOMIZATION",
		},
		COMMUNAL: ListScheduledTaskRequestAssociatedTaskType{
			value: "COMMUNAL",
		},
	}
}

func (c ListScheduledTaskRequestAssociatedTaskType) Value() string {
	return c.value
}

func (c ListScheduledTaskRequestAssociatedTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskRequestAssociatedTaskType) UnmarshalJSON(b []byte) error {
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

type ListScheduledTaskRequestApproveStatus struct {
	value string
}

type ListScheduledTaskRequestApproveStatusEnum struct {
	PENDING  ListScheduledTaskRequestApproveStatus
	REJECTED ListScheduledTaskRequestApproveStatus
	PASSED   ListScheduledTaskRequestApproveStatus
}

func GetListScheduledTaskRequestApproveStatusEnum() ListScheduledTaskRequestApproveStatusEnum {
	return ListScheduledTaskRequestApproveStatusEnum{
		PENDING: ListScheduledTaskRequestApproveStatus{
			value: "PENDING",
		},
		REJECTED: ListScheduledTaskRequestApproveStatus{
			value: "REJECTED",
		},
		PASSED: ListScheduledTaskRequestApproveStatus{
			value: "PASSED",
		},
	}
}

func (c ListScheduledTaskRequestApproveStatus) Value() string {
	return c.value
}

func (c ListScheduledTaskRequestApproveStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskRequestApproveStatus) UnmarshalJSON(b []byte) error {
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

type ListScheduledTaskRequestLastExecutionStatus struct {
	value string
}

type ListScheduledTaskRequestLastExecutionStatusEnum struct {
	READY      ListScheduledTaskRequestLastExecutionStatus
	PROCESSING ListScheduledTaskRequestLastExecutionStatus
	ABNORMAL   ListScheduledTaskRequestLastExecutionStatus
	FINISHED   ListScheduledTaskRequestLastExecutionStatus
	PAUSED     ListScheduledTaskRequestLastExecutionStatus
	CANCELED   ListScheduledTaskRequestLastExecutionStatus
}

func GetListScheduledTaskRequestLastExecutionStatusEnum() ListScheduledTaskRequestLastExecutionStatusEnum {
	return ListScheduledTaskRequestLastExecutionStatusEnum{
		READY: ListScheduledTaskRequestLastExecutionStatus{
			value: "READY",
		},
		PROCESSING: ListScheduledTaskRequestLastExecutionStatus{
			value: "PROCESSING",
		},
		ABNORMAL: ListScheduledTaskRequestLastExecutionStatus{
			value: "ABNORMAL",
		},
		FINISHED: ListScheduledTaskRequestLastExecutionStatus{
			value: "FINISHED",
		},
		PAUSED: ListScheduledTaskRequestLastExecutionStatus{
			value: "PAUSED",
		},
		CANCELED: ListScheduledTaskRequestLastExecutionStatus{
			value: "CANCELED",
		},
	}
}

func (c ListScheduledTaskRequestLastExecutionStatus) Value() string {
	return c.value
}

func (c ListScheduledTaskRequestLastExecutionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListScheduledTaskRequestLastExecutionStatus) UnmarshalJSON(b []byte) error {
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
