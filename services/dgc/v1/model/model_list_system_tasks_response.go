package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type ListSystemTasksResponse struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	StartTime *string `json:"startTime,omitempty"`

	EndTime *string `json:"endTime,omitempty"`

	LastUpdate *string `json:"lastUpdate,omitempty"`

	Status *ListSystemTasksResponseStatus `json:"status,omitempty"`

	Message *string `json:"message,omitempty"`

	SubTasks       *[]SubTaskStatus `json:"subTasks,omitempty"`
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

func (c ListSystemTasksResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSystemTasksResponseStatus) UnmarshalJSON(b []byte) error {
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
