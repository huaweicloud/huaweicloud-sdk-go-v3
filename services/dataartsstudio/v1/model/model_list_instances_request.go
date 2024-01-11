package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListInstancesRequest Request Object
type ListInstancesRequest struct {

	// DataArts Studio工作空间ID
	Workspace string `json:"workspace"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 任务实例类型 QUALITY_TASK:质量作业 CONSISTENCY_TASK:对账作业
	TaskType *ListInstancesRequestTaskType `json:"task_type,omitempty"`

	// 状态, RUNNING:运行中,FAILED:失败,ALARMING:报警,SUCCESS:正常,SUSPENDING:暂停中,UNKNOWN:未定义
	RunStatus *ListInstancesRequestRunStatus `json:"run_status,omitempty"`

	// 通知状态 NOT_TRIGGERED:未触发,SUCCESS:成功,FAILED:失败
	NotifyStatus *ListInstancesRequestNotifyStatus `json:"notify_status,omitempty"`

	// 最近运行时间查询区间的开始时间,13位时间戳(精确到毫秒)
	StartTime *int64 `json:"start_time,omitempty"`

	// 最近运行时间查询区间的结束时间,13位时间戳(精确到毫秒)
	EndTime *int64 `json:"end_time,omitempty"`

	// 每页显示的条目数量,最大值为100
	Limit *int32 `json:"limit,omitempty"`

	// 分页偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestTaskType struct {
	value string
}

type ListInstancesRequestTaskTypeEnum struct {
	QUALITY_TASK     ListInstancesRequestTaskType
	CONSISTENCY_TASK ListInstancesRequestTaskType
}

func GetListInstancesRequestTaskTypeEnum() ListInstancesRequestTaskTypeEnum {
	return ListInstancesRequestTaskTypeEnum{
		QUALITY_TASK: ListInstancesRequestTaskType{
			value: "QUALITY_TASK",
		},
		CONSISTENCY_TASK: ListInstancesRequestTaskType{
			value: "CONSISTENCY_TASK",
		},
	}
}

func (c ListInstancesRequestTaskType) Value() string {
	return c.value
}

func (c ListInstancesRequestTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestTaskType) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestRunStatus struct {
	value string
}

type ListInstancesRequestRunStatusEnum struct {
	SUCCESS  ListInstancesRequestRunStatus
	FAILED   ListInstancesRequestRunStatus
	RUNNING  ListInstancesRequestRunStatus
	ALARMING ListInstancesRequestRunStatus
}

func GetListInstancesRequestRunStatusEnum() ListInstancesRequestRunStatusEnum {
	return ListInstancesRequestRunStatusEnum{
		SUCCESS: ListInstancesRequestRunStatus{
			value: "SUCCESS",
		},
		FAILED: ListInstancesRequestRunStatus{
			value: "FAILED",
		},
		RUNNING: ListInstancesRequestRunStatus{
			value: "RUNNING",
		},
		ALARMING: ListInstancesRequestRunStatus{
			value: "ALARMING",
		},
	}
}

func (c ListInstancesRequestRunStatus) Value() string {
	return c.value
}

func (c ListInstancesRequestRunStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestRunStatus) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestNotifyStatus struct {
	value string
}

type ListInstancesRequestNotifyStatusEnum struct {
	SUCCESS       ListInstancesRequestNotifyStatus
	FAILED        ListInstancesRequestNotifyStatus
	NOT_TRIGGERED ListInstancesRequestNotifyStatus
}

func GetListInstancesRequestNotifyStatusEnum() ListInstancesRequestNotifyStatusEnum {
	return ListInstancesRequestNotifyStatusEnum{
		SUCCESS: ListInstancesRequestNotifyStatus{
			value: "SUCCESS",
		},
		FAILED: ListInstancesRequestNotifyStatus{
			value: "FAILED",
		},
		NOT_TRIGGERED: ListInstancesRequestNotifyStatus{
			value: "NOT_TRIGGERED",
		},
	}
}

func (c ListInstancesRequestNotifyStatus) Value() string {
	return c.value
}

func (c ListInstancesRequestNotifyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListInstancesRequestNotifyStatus) UnmarshalJSON(b []byte) error {
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
