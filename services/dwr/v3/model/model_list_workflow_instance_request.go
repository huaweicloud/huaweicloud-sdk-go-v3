package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListWorkflowInstanceRequest Request Object
type ListWorkflowInstanceRequest struct {

	// 请求返回的最大记录条数。分页查询，每页显示的条目数量，最大数量200，超过200后只返回200
	Limit *int32 `json:"limit,omitempty"`

	// 工作流名称。
	GraphName string `json:"graph_name"`

	// 查询开始时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间。若起始时间未填写，以终止时间前推3天为起始时间
	StartTime *string `json:"start_time,omitempty"`

	// 查询终止时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间。若终止时间未填写，以起始时间后退3天未终止时间。若均未填写，默认查询最近3天数据。
	EndTime *string `json:"end_time,omitempty"`

	// 需要过滤的流程实例状态  最小长度：0  最大长度：64  枚举值：  success  fail  running  timeout  cancel
	Status *ListWorkflowInstanceRequestStatus `json:"status,omitempty"`

	// 查询的起始位置。start大于等于1，最大1000，不设置则取默认值1。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWorkflowInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowInstanceRequest", string(data)}, " ")
}

type ListWorkflowInstanceRequestStatus struct {
	value string
}

type ListWorkflowInstanceRequestStatusEnum struct {
	SUCCESS ListWorkflowInstanceRequestStatus
	FAIL    ListWorkflowInstanceRequestStatus
	RUNNING ListWorkflowInstanceRequestStatus
	TIMEOUT ListWorkflowInstanceRequestStatus
	CANCEL  ListWorkflowInstanceRequestStatus
}

func GetListWorkflowInstanceRequestStatusEnum() ListWorkflowInstanceRequestStatusEnum {
	return ListWorkflowInstanceRequestStatusEnum{
		SUCCESS: ListWorkflowInstanceRequestStatus{
			value: "success",
		},
		FAIL: ListWorkflowInstanceRequestStatus{
			value: "fail",
		},
		RUNNING: ListWorkflowInstanceRequestStatus{
			value: "running",
		},
		TIMEOUT: ListWorkflowInstanceRequestStatus{
			value: "timeout",
		},
		CANCEL: ListWorkflowInstanceRequestStatus{
			value: "cancel",
		},
	}
}

func (c ListWorkflowInstanceRequestStatus) Value() string {
	return c.value
}

func (c ListWorkflowInstanceRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowInstanceRequestStatus) UnmarshalJSON(b []byte) error {
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
