package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Instance struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例状态。
	Status *InstanceStatus `json:"status,omitempty"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}

type InstanceStatus struct {
	value string
}

type InstanceStatusEnum struct {
	CREATING    InstanceStatus
	RUNNING     InstanceStatus
	ABNORMAL    InstanceStatus
	FAILED      InstanceStatus
	TERMINATING InstanceStatus
}

func GetInstanceStatusEnum() InstanceStatusEnum {
	return InstanceStatusEnum{
		CREATING: InstanceStatus{
			value: "creating",
		},
		RUNNING: InstanceStatus{
			value: "running",
		},
		ABNORMAL: InstanceStatus{
			value: "abnormal",
		},
		FAILED: InstanceStatus{
			value: "failed",
		},
		TERMINATING: InstanceStatus{
			value: "terminating",
		},
	}
}

func (c InstanceStatus) Value() string {
	return c.value
}

func (c InstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceStatus) UnmarshalJSON(b []byte) error {
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
