package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowExecutionsRequest Request Object
type ListWorkflowExecutionsRequest struct {

	// 函数工作流ID
	WorkflowId string `json:"workflow_id"`

	// 分页查询，每页显示的条目数量，最大数量200，超过200后只返回200
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询，分页的偏移量，默认值为0 offset小于0时，按照0处理
	Offset *int32 `json:"offset,omitempty"`

	// 需要过滤的流程实例状态
	Status *ListWorkflowExecutionsRequestStatus `json:"status,omitempty"`

	// 查询开始时间，UTC时间。若起始时间未填写，以终止时间前推3天为起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 查询开始时间，UTC时间。若终止时间未填写，以起始时间后退3天未终止时间。若均未填写，默认查询最近3天数据。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListWorkflowExecutionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowExecutionsRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowExecutionsRequest", string(data)}, " ")
}

type ListWorkflowExecutionsRequestStatus struct {
	value string
}

type ListWorkflowExecutionsRequestStatusEnum struct {
	SUCCESS ListWorkflowExecutionsRequestStatus
	FAIL    ListWorkflowExecutionsRequestStatus
	RUNNING ListWorkflowExecutionsRequestStatus
	TIMEOUT ListWorkflowExecutionsRequestStatus
	CANCEL  ListWorkflowExecutionsRequestStatus
}

func GetListWorkflowExecutionsRequestStatusEnum() ListWorkflowExecutionsRequestStatusEnum {
	return ListWorkflowExecutionsRequestStatusEnum{
		SUCCESS: ListWorkflowExecutionsRequestStatus{
			value: "success",
		},
		FAIL: ListWorkflowExecutionsRequestStatus{
			value: "fail",
		},
		RUNNING: ListWorkflowExecutionsRequestStatus{
			value: "running",
		},
		TIMEOUT: ListWorkflowExecutionsRequestStatus{
			value: "timeout",
		},
		CANCEL: ListWorkflowExecutionsRequestStatus{
			value: "cancel",
		},
	}
}

func (c ListWorkflowExecutionsRequestStatus) Value() string {
	return c.value
}

func (c ListWorkflowExecutionsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowExecutionsRequestStatus) UnmarshalJSON(b []byte) error {
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
