package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListHotKeyScanTasksRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 分析任务状态
	Status *ListHotKeyScanTasksRequestStatus `json:"status,omitempty"`
}

func (o ListHotKeyScanTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotKeyScanTasksRequest struct{}"
	}

	return strings.Join([]string{"ListHotKeyScanTasksRequest", string(data)}, " ")
}

type ListHotKeyScanTasksRequestStatus struct {
	value string
}

type ListHotKeyScanTasksRequestStatusEnum struct {
	WAITING ListHotKeyScanTasksRequestStatus
	RUNNING ListHotKeyScanTasksRequestStatus
	SUCCESS ListHotKeyScanTasksRequestStatus
	FAILED  ListHotKeyScanTasksRequestStatus
}

func GetListHotKeyScanTasksRequestStatusEnum() ListHotKeyScanTasksRequestStatusEnum {
	return ListHotKeyScanTasksRequestStatusEnum{
		WAITING: ListHotKeyScanTasksRequestStatus{
			value: "waiting",
		},
		RUNNING: ListHotKeyScanTasksRequestStatus{
			value: "running",
		},
		SUCCESS: ListHotKeyScanTasksRequestStatus{
			value: "success",
		},
		FAILED: ListHotKeyScanTasksRequestStatus{
			value: "failed",
		},
	}
}

func (c ListHotKeyScanTasksRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListHotKeyScanTasksRequestStatus) UnmarshalJSON(b []byte) error {
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
