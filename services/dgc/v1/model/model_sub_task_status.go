package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type SubTaskStatus struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	StartTime *string `json:"startTime,omitempty"`

	EndTime *string `json:"endTime,omitempty"`

	LastUpdate *string `json:"lastUpdate,omitempty"`

	Status *SubTaskStatusStatus `json:"status,omitempty"`
}

func (o SubTaskStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskStatus struct{}"
	}

	return strings.Join([]string{"SubTaskStatus", string(data)}, " ")
}

type SubTaskStatusStatus struct {
	value string
}

type SubTaskStatusStatusEnum struct {
	RUNNING    SubTaskStatusStatus
	SUCCESSFUL SubTaskStatusStatus
	FAILED     SubTaskStatusStatus
}

func GetSubTaskStatusStatusEnum() SubTaskStatusStatusEnum {
	return SubTaskStatusStatusEnum{
		RUNNING: SubTaskStatusStatus{
			value: "RUNNING",
		},
		SUCCESSFUL: SubTaskStatusStatus{
			value: "SUCCESSFUL",
		},
		FAILED: SubTaskStatusStatus{
			value: "FAILED",
		},
	}
}

func (c SubTaskStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubTaskStatusStatus) UnmarshalJSON(b []byte) error {
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
