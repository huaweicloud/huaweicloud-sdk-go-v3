package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourceLimit struct {

	// cpu限额。
	CpuLimit ResourceLimitCpuLimit `json:"cpu_limit"`

	// 内存限额。
	MemoryLimit ResourceLimitMemoryLimit `json:"memory_limit"`
}

func (o ResourceLimit) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceLimit struct{}"
	}

	return strings.Join([]string{"ResourceLimit", string(data)}, " ")
}

type ResourceLimitCpuLimit struct {
	value string
}

type ResourceLimitCpuLimitEnum struct {
	E_500M  ResourceLimitCpuLimit
	E_1000M ResourceLimitCpuLimit
	E_2000M ResourceLimitCpuLimit
}

func GetResourceLimitCpuLimitEnum() ResourceLimitCpuLimitEnum {
	return ResourceLimitCpuLimitEnum{
		E_500M: ResourceLimitCpuLimit{
			value: "500m",
		},
		E_1000M: ResourceLimitCpuLimit{
			value: "1000m",
		},
		E_2000M: ResourceLimitCpuLimit{
			value: "2000m",
		},
	}
}

func (c ResourceLimitCpuLimit) Value() string {
	return c.value
}

func (c ResourceLimitCpuLimit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceLimitCpuLimit) UnmarshalJSON(b []byte) error {
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

type ResourceLimitMemoryLimit struct {
	value string
}

type ResourceLimitMemoryLimitEnum struct {
	E_1_GI ResourceLimitMemoryLimit
	E_2_GI ResourceLimitMemoryLimit
	E_4_GI ResourceLimitMemoryLimit
}

func GetResourceLimitMemoryLimitEnum() ResourceLimitMemoryLimitEnum {
	return ResourceLimitMemoryLimitEnum{
		E_1_GI: ResourceLimitMemoryLimit{
			value: "1Gi",
		},
		E_2_GI: ResourceLimitMemoryLimit{
			value: "2Gi",
		},
		E_4_GI: ResourceLimitMemoryLimit{
			value: "4Gi",
		},
	}
}

func (c ResourceLimitMemoryLimit) Value() string {
	return c.value
}

func (c ResourceLimitMemoryLimit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourceLimitMemoryLimit) UnmarshalJSON(b []byte) error {
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
