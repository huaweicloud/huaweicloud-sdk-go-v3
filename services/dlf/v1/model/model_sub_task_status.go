package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

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

func (c SubTaskStatusStatus) Value() string {
	return c.value
}

func (c SubTaskStatusStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubTaskStatusStatus) UnmarshalJSON(b []byte) error {
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
