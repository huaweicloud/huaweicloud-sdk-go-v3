package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 返回值
type InstancesResponseInstancesVoResult struct {
	// 链接

	Link *string `json:"link,omitempty"`
	// cpu架构 x86|arm

	Arch *InstancesResponseInstancesVoResultArch `json:"arch,omitempty"`
	// 实例id

	Id *string `json:"id,omitempty"`
	// 是否私有平台

	Private *bool `json:"private,omitempty"`
}

func (o InstancesResponseInstancesVoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesResponseInstancesVoResult struct{}"
	}

	return strings.Join([]string{"InstancesResponseInstancesVoResult", string(data)}, " ")
}

type InstancesResponseInstancesVoResultArch struct {
	value string
}

type InstancesResponseInstancesVoResultArchEnum struct {
	X86 InstancesResponseInstancesVoResultArch
	ARM InstancesResponseInstancesVoResultArch
}

func GetInstancesResponseInstancesVoResultArchEnum() InstancesResponseInstancesVoResultArchEnum {
	return InstancesResponseInstancesVoResultArchEnum{
		X86: InstancesResponseInstancesVoResultArch{
			value: "x86",
		},
		ARM: InstancesResponseInstancesVoResultArch{
			value: "arm",
		},
	}
}

func (c InstancesResponseInstancesVoResultArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancesResponseInstancesVoResultArch) UnmarshalJSON(b []byte) error {
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
