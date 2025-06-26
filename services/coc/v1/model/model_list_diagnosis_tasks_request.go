package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDiagnosisTasksRequest Request Object
type ListDiagnosisTasksRequest struct {

	// 诊断任务工单ID
	TaskId *string `json:"task_id,omitempty"`

	// 诊断任务实例类别
	Type *ListDiagnosisTasksRequestType `json:"type,omitempty"`

	// 诊断任务执行状态
	Status *ListDiagnosisTasksRequestStatus `json:"status,omitempty"`

	// 被诊断实例所在区域
	Region *string `json:"region,omitempty"`

	// 诊断工单创建者
	Creator *string `json:"creator,omitempty"`

	// 诊断工单的开始执行时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 诊断工单的执行结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 分页查询页索引
	PageIndex int32 `json:"page_index"`

	// 分页查询页大小
	PageSize int32 `json:"page_size"`
}

func (o ListDiagnosisTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDiagnosisTasksRequest struct{}"
	}

	return strings.Join([]string{"ListDiagnosisTasksRequest", string(data)}, " ")
}

type ListDiagnosisTasksRequestType struct {
	value string
}

type ListDiagnosisTasksRequestTypeEnum struct {
	ECS ListDiagnosisTasksRequestType
	RDS ListDiagnosisTasksRequestType
	DCS ListDiagnosisTasksRequestType
	DMS ListDiagnosisTasksRequestType
	ELB ListDiagnosisTasksRequestType
}

func GetListDiagnosisTasksRequestTypeEnum() ListDiagnosisTasksRequestTypeEnum {
	return ListDiagnosisTasksRequestTypeEnum{
		ECS: ListDiagnosisTasksRequestType{
			value: "ECS",
		},
		RDS: ListDiagnosisTasksRequestType{
			value: "RDS",
		},
		DCS: ListDiagnosisTasksRequestType{
			value: "DCS",
		},
		DMS: ListDiagnosisTasksRequestType{
			value: "DMS",
		},
		ELB: ListDiagnosisTasksRequestType{
			value: "ELB",
		},
	}
}

func (c ListDiagnosisTasksRequestType) Value() string {
	return c.value
}

func (c ListDiagnosisTasksRequestType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDiagnosisTasksRequestType) UnmarshalJSON(b []byte) error {
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

type ListDiagnosisTasksRequestStatus struct {
	value string
}

type ListDiagnosisTasksRequestStatusEnum struct {
	CANCEL    ListDiagnosisTasksRequestStatus
	EXECUTING ListDiagnosisTasksRequestStatus
	WAITING   ListDiagnosisTasksRequestStatus
	FAILED    ListDiagnosisTasksRequestStatus
	FINISH    ListDiagnosisTasksRequestStatus
}

func GetListDiagnosisTasksRequestStatusEnum() ListDiagnosisTasksRequestStatusEnum {
	return ListDiagnosisTasksRequestStatusEnum{
		CANCEL: ListDiagnosisTasksRequestStatus{
			value: "cancel",
		},
		EXECUTING: ListDiagnosisTasksRequestStatus{
			value: "executing",
		},
		WAITING: ListDiagnosisTasksRequestStatus{
			value: "waiting",
		},
		FAILED: ListDiagnosisTasksRequestStatus{
			value: "failed",
		},
		FINISH: ListDiagnosisTasksRequestStatus{
			value: "finish",
		},
	}
}

func (c ListDiagnosisTasksRequestStatus) Value() string {
	return c.value
}

func (c ListDiagnosisTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDiagnosisTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
