package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobNodeBaseInfo 任务实例基本信息体。
type JobNodeBaseInfo struct {

	// 实例类型。取值： - single：单机。 - ha：主备。
	InstanceType JobNodeBaseInfoInstanceType `json:"instance_type"`

	// CPU架构。取值： - x86 - arm
	Arch JobNodeBaseInfoArch `json:"arch"`

	// 可用区ID。 约束：对于任务实例类型不是单机的实例，需要分别为实例所有节点指定可用区，并用“,”英文逗号隔开。示例： - 实例类型为single：\"cn-north-4a\" - 实例类型为ha：\"cn-north-4a,cn-north-4b\"
	AvailabilityZone string `json:"availability_zone"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 任务主备角色。
	Role *string `json:"role,omitempty"`
}

func (o JobNodeBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobNodeBaseInfo struct{}"
	}

	return strings.Join([]string{"JobNodeBaseInfo", string(data)}, " ")
}

type JobNodeBaseInfoInstanceType struct {
	value string
}

type JobNodeBaseInfoInstanceTypeEnum struct {
	SINGLE JobNodeBaseInfoInstanceType
	HA     JobNodeBaseInfoInstanceType
}

func GetJobNodeBaseInfoInstanceTypeEnum() JobNodeBaseInfoInstanceTypeEnum {
	return JobNodeBaseInfoInstanceTypeEnum{
		SINGLE: JobNodeBaseInfoInstanceType{
			value: "single",
		},
		HA: JobNodeBaseInfoInstanceType{
			value: "ha",
		},
	}
}

func (c JobNodeBaseInfoInstanceType) Value() string {
	return c.value
}

func (c JobNodeBaseInfoInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobNodeBaseInfoInstanceType) UnmarshalJSON(b []byte) error {
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

type JobNodeBaseInfoArch struct {
	value string
}

type JobNodeBaseInfoArchEnum struct {
	X86 JobNodeBaseInfoArch
	ARM JobNodeBaseInfoArch
}

func GetJobNodeBaseInfoArchEnum() JobNodeBaseInfoArchEnum {
	return JobNodeBaseInfoArchEnum{
		X86: JobNodeBaseInfoArch{
			value: "x86",
		},
		ARM: JobNodeBaseInfoArch{
			value: "arm",
		},
	}
}

func (c JobNodeBaseInfoArch) Value() string {
	return c.value
}

func (c JobNodeBaseInfoArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobNodeBaseInfoArch) UnmarshalJSON(b []byte) error {
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
