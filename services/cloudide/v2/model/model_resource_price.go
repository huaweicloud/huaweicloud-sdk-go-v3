package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourcePrice struct {
	// cpu架构 x86|arm

	Arch *ResourcePriceArch `json:"arch,omitempty"`
	// 价格

	Price *float32 `json:"price,omitempty"`
	// 规格。 类型为'storage'时，size值可以为5GB，10GB，20GB。 类型为'cpuMemory'时，arch为'x86'，size值可以为1U1G，2U4G；arch为'arm'，size值可以为4U8G。

	Size *string `json:"size,omitempty"`
	// 类型。目前可以取值storage，cpuMemory

	Type *string `json:"type,omitempty"`
}

func (o ResourcePrice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePrice struct{}"
	}

	return strings.Join([]string{"ResourcePrice", string(data)}, " ")
}

type ResourcePriceArch struct {
	value string
}

type ResourcePriceArchEnum struct {
	X86 ResourcePriceArch
	ARM ResourcePriceArch
}

func GetResourcePriceArchEnum() ResourcePriceArchEnum {
	return ResourcePriceArchEnum{
		X86: ResourcePriceArch{
			value: "x86",
		},
		ARM: ResourcePriceArch{
			value: "arm",
		},
	}
}

func (c ResourcePriceArch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcePriceArch) UnmarshalJSON(b []byte) error {
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
