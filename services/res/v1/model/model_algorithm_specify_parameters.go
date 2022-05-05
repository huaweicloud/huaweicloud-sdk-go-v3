package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

//
type AlgorithmSpecifyParameters struct {

	// 隐向量长度(DEEPFM需要提供此参数)。
	LatentVectorLength *int32 `json:"latent_vector_length,omitempty"`

	// 神经网络结构(DEEPFM需要提供此参数)。
	Architecture *[]int32 `json:"architecture,omitempty"`

	// 激活函数(DEEPFM需要提供此参数,AutoGroup需要提供此参数)。
	ActiveFunction *AlgorithmSpecifyParametersActiveFunction `json:"active_function,omitempty"`

	// 神经元值保留概率(DEEPFM需要提供此参数,AutoGroup需要提供此参数)。
	ValueKeepProbability *float64 `json:"value_keep_probability,omitempty"`

	// 各阶隐向量长度(AutoGroup需要提供此参数)。
	EmbedSize *[]int32 `json:"embed_size,omitempty"`

	// 神经网络结构(AutoGroup需要提供此参数)。
	MlpArchitecture *[]int32 `json:"mlp_architecture,omitempty"`

	// 最大交互阶数(AutoGroup需要提供此参数)。
	MaxOrder *int32 `json:"max_order,omitempty"`

	// 哈希长度(AutoGroup需要提供此参数)。
	HashSizes *[]int32 `json:"hash_sizes,omitempty"`

	// 特征交互层惩罚项系数(AutoGroup需要提供此参数)。
	HashCompensation *[]float64 `json:"hash_compensation,omitempty"`

	// 使用线性部分(AutoGroup需要提供此参数)。
	UseWidePart *bool `json:"use_wide_part,omitempty"`

	StructureOptimizer *Optimizer `json:"structure_optimizer,omitempty"`

	// 融合多值特征(AutoGroup需要提供此参数)。
	MergeMultiHot *bool `json:"merge_multi_hot,omitempty"`

	// 固定哈希结构(AutoGroup需要提供此参数)。
	FixStructure *bool `json:"fix_structure,omitempty"`
}

func (o AlgorithmSpecifyParameters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmSpecifyParameters struct{}"
	}

	return strings.Join([]string{"AlgorithmSpecifyParameters", string(data)}, " ")
}

type AlgorithmSpecifyParametersActiveFunction struct {
	value string
}

type AlgorithmSpecifyParametersActiveFunctionEnum struct {
	RELU    AlgorithmSpecifyParametersActiveFunction
	SIGMOID AlgorithmSpecifyParametersActiveFunction
	TANH    AlgorithmSpecifyParametersActiveFunction
}

func GetAlgorithmSpecifyParametersActiveFunctionEnum() AlgorithmSpecifyParametersActiveFunctionEnum {
	return AlgorithmSpecifyParametersActiveFunctionEnum{
		RELU: AlgorithmSpecifyParametersActiveFunction{
			value: "relu",
		},
		SIGMOID: AlgorithmSpecifyParametersActiveFunction{
			value: "sigmoid",
		},
		TANH: AlgorithmSpecifyParametersActiveFunction{
			value: "tanh",
		},
	}
}

func (c AlgorithmSpecifyParametersActiveFunction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlgorithmSpecifyParametersActiveFunction) UnmarshalJSON(b []byte) error {
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
