package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListSystemTasksResponse Response Object
type ListSystemTasksResponse struct {

	// 作业ID
	Id *string `json:"id,omitempty"`

	// 作业名称
	Name *string `json:"name,omitempty"`

	// 作业开始日期 13位时间戳
	StartTime *int64 `json:"startTime,omitempty"`

	// 作业结束日期 13位时间戳
	EndTime *int64 `json:"endTime,omitempty"`

	// 作业最后更新日期 13位时间戳
	LastUpdate *int64 `json:"lastUpdate,omitempty"`

	// 作业运行状态 RUNNING：运行中 SUCCESSFUL：运行成功 FAILED：运行失败
	Status *ListSystemTasksResponseStatus `json:"status,omitempty"`

	// 作业消息
	Message *string `json:"message,omitempty"`

	// 当前作业包含的子作业
	Subtasks       *[]SubTaskStatus `json:"subtasks,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListSystemTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSystemTasksResponse struct{}"
	}

	return strings.Join([]string{"ListSystemTasksResponse", string(data)}, " ")
}

type ListSystemTasksResponseStatus struct {
	value string
}

type ListSystemTasksResponseStatusEnum struct {
	RUNNING    ListSystemTasksResponseStatus
	SUCCESSFUL ListSystemTasksResponseStatus
	FAILED     ListSystemTasksResponseStatus
}

func GetListSystemTasksResponseStatusEnum() ListSystemTasksResponseStatusEnum {
	return ListSystemTasksResponseStatusEnum{
		RUNNING: ListSystemTasksResponseStatus{
			value: "RUNNING",
		},
		SUCCESSFUL: ListSystemTasksResponseStatus{
			value: "SUCCESSFUL",
		},
		FAILED: ListSystemTasksResponseStatus{
			value: "FAILED",
		},
	}
}

func (c ListSystemTasksResponseStatus) Value() string {
	return c.value
}

func (c ListSystemTasksResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSystemTasksResponseStatus) UnmarshalJSON(b []byte) error {
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
