package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ComponentProbe struct {
	Type ComponentProbeType `json:"type"`

	// 表示启动后多久开始探测
	Delay *int32 `json:"delay,omitempty"`

	// 表示探测超时时间
	Timeout *int32 `json:"timeout,omitempty"`

	Parameters *ProbeParameter `json:"parameters"`
}

func (o ComponentProbe) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentProbe struct{}"
	}

	return strings.Join([]string{"ComponentProbe", string(data)}, " ")
}

type ComponentProbeType struct {
	value string
}

type ComponentProbeTypeEnum struct {
	HTTP    ComponentProbeType
	TCP     ComponentProbeType
	COMMAND ComponentProbeType
}

func GetComponentProbeTypeEnum() ComponentProbeTypeEnum {
	return ComponentProbeTypeEnum{
		HTTP: ComponentProbeType{
			value: "http",
		},
		TCP: ComponentProbeType{
			value: "tcp",
		},
		COMMAND: ComponentProbeType{
			value: "command",
		},
	}
}

func (c ComponentProbeType) Value() string {
	return c.value
}

func (c ComponentProbeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ComponentProbeType) UnmarshalJSON(b []byte) error {
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
