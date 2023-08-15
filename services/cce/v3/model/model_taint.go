package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Taint 如下字段不可使用：  - node.kubernetes.io/memory-pressure - node.kubernetes.io/disk-pressure - node.kubernetes.io/out-of-disk - node.kubernetes.io/unschedulable - node.kubernetes.io/network-unavailable
type Taint struct {

	// 键
	Key string `json:"key"`

	// 值
	Value *string `json:"value,omitempty"`

	// 作用效果
	Effect TaintEffect `json:"effect"`
}

func (o Taint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Taint struct{}"
	}

	return strings.Join([]string{"Taint", string(data)}, " ")
}

type TaintEffect struct {
	value string
}

type TaintEffectEnum struct {
	NO_SCHEDULE        TaintEffect
	PREFER_NO_SCHEDULE TaintEffect
	NO_EXECUTE         TaintEffect
}

func GetTaintEffectEnum() TaintEffectEnum {
	return TaintEffectEnum{
		NO_SCHEDULE: TaintEffect{
			value: "NoSchedule",
		},
		PREFER_NO_SCHEDULE: TaintEffect{
			value: "PreferNoSchedule",
		},
		NO_EXECUTE: TaintEffect{
			value: "NoExecute",
		},
	}
}

func (c TaintEffect) Value() string {
	return c.value
}

func (c TaintEffect) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaintEffect) UnmarshalJSON(b []byte) error {
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
