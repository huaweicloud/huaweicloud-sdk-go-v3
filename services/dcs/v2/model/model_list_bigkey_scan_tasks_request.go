package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListBigkeyScanTasksRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 偏移量，表示从此偏移量开始查询， offset大于等于0

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的条目数量。

	Limit *int32 `json:"limit,omitempty"`
	// 分析任务状态

	Status *ListBigkeyScanTasksRequestStatus `json:"status,omitempty"`
}

func (o ListBigkeyScanTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBigkeyScanTasksRequest struct{}"
	}

	return strings.Join([]string{"ListBigkeyScanTasksRequest", string(data)}, " ")
}

type ListBigkeyScanTasksRequestStatus struct {
	value string
}

type ListBigkeyScanTasksRequestStatusEnum struct {
	WAITING ListBigkeyScanTasksRequestStatus
	RUNNING ListBigkeyScanTasksRequestStatus
	SUCCESS ListBigkeyScanTasksRequestStatus
	FAILED  ListBigkeyScanTasksRequestStatus
}

func GetListBigkeyScanTasksRequestStatusEnum() ListBigkeyScanTasksRequestStatusEnum {
	return ListBigkeyScanTasksRequestStatusEnum{
		WAITING: ListBigkeyScanTasksRequestStatus{
			value: "waiting",
		},
		RUNNING: ListBigkeyScanTasksRequestStatus{
			value: "running",
		},
		SUCCESS: ListBigkeyScanTasksRequestStatus{
			value: "success",
		},
		FAILED: ListBigkeyScanTasksRequestStatus{
			value: "failed",
		},
	}
}

func (c ListBigkeyScanTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBigkeyScanTasksRequestStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
