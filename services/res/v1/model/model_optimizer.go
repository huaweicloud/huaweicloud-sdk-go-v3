package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Optimizer 优化器参数
type Optimizer struct {

	// 优化器类型。
	Type *OptimizerType `json:"type,omitempty"`

	// 学习率。
	LearningRate *float64 `json:"learning_rate,omitempty"`

	// 初始梯度累加和。
	InitialAccumulatorValue *float64 `json:"initial_accumulator_value,omitempty"`

	// L1正则项系数。
	Lambda1 *float64 `json:"lambda1,omitempty"`

	// L2正则项系数。
	Lambda2 *float64 `json:"lambda2,omitempty"`

	// 数值稳定常量。
	Epsilon *float64 `json:"epsilon,omitempty"`

	// 衰减因子。
	DecayRate *float64 `json:"decay_rate,omitempty"`

	// 衰减步长。
	DecaySteps *float64 `json:"decay_steps,omitempty"`
}

func (o Optimizer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Optimizer struct{}"
	}

	return strings.Join([]string{"Optimizer", string(data)}, " ")
}

type OptimizerType struct {
	value string
}

type OptimizerTypeEnum struct {
	ADAM    OptimizerType
	ADAGRAD OptimizerType
	FTRL    OptimizerType
}

func GetOptimizerTypeEnum() OptimizerTypeEnum {
	return OptimizerTypeEnum{
		ADAM: OptimizerType{
			value: "adam",
		},
		ADAGRAD: OptimizerType{
			value: "adagrad",
		},
		FTRL: OptimizerType{
			value: "ftrl",
		},
	}
}

func (c OptimizerType) Value() string {
	return c.value
}

func (c OptimizerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OptimizerType) UnmarshalJSON(b []byte) error {
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
