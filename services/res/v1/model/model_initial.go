package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Initial 初始化参数
type Initial struct {

	// 初始化方法。
	InitialMethod InitialInitialMethod `json:"initial_method"`

	// 平均值。
	MeanValue *float64 `json:"mean_value,omitempty"`

	// 标准差。
	StandardDeviation *float64 `json:"standard_deviation,omitempty"`

	// 最小值。
	MinValue *float64 `json:"min_value,omitempty"`

	// 最大值。
	MaxValue *float64 `json:"max_value,omitempty"`
}

func (o Initial) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Initial struct{}"
	}

	return strings.Join([]string{"Initial", string(data)}, " ")
}

type InitialInitialMethod struct {
	value string
}

type InitialInitialMethodEnum struct {
	NORMAL  InitialInitialMethod
	UNIFORM InitialInitialMethod
	XAVIER  InitialInitialMethod
}

func GetInitialInitialMethodEnum() InitialInitialMethodEnum {
	return InitialInitialMethodEnum{
		NORMAL: InitialInitialMethod{
			value: "normal",
		},
		UNIFORM: InitialInitialMethod{
			value: "uniform",
		},
		XAVIER: InitialInitialMethod{
			value: "xavier",
		},
	}
}

func (c InitialInitialMethod) Value() string {
	return c.value
}

func (c InitialInitialMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InitialInitialMethod) UnmarshalJSON(b []byte) error {
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
