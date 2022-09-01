package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ListSystemTasksResponse struct {
	Id *string `json:"id,omitempty" xml:"id"`

	Name *string `json:"name,omitempty" xml:"name"`

	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	EndTime *string `json:"endTime,omitempty" xml:"endTime"`

	LastUpdate *string `json:"lastUpdate,omitempty" xml:"lastUpdate"`

	Status *ListSystemTasksResponseStatus `json:"status,omitempty" xml:"status"`

	Message *string `json:"message,omitempty" xml:"message"`

	SubTasks       *[]SubTaskStatus `json:"subTasks,omitempty" xml:"subTasks"`
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
