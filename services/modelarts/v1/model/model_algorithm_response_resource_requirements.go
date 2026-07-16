package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlgorithmResponseResourceRequirements struct {

	// 资源约束，可选值如下： - 资源类型（flavor_type），对应值可选择CPU、GPU[或Ascend](tag:hc,hk,fcs_super)； - 是否支持多卡训练（device_distributed_mode），对应值可选择支持（multiple）、不支持（singular）； - 是否支持分布式训练（host_distributed_mode），对应值可选择支持（multiple）、不支持（singular）。
	Key *string `json:"key,omitempty"`

	// 资源约束键对应值。
	Value *[]string `json:"value,omitempty"`

	// 键与值关系，当前只支持in。例如flavor_type in [CPU,GPU]。
	Operator *string `json:"operator,omitempty"`
}

func (o AlgorithmResponseResourceRequirements) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponseResourceRequirements struct{}"
	}

	return strings.Join([]string{"AlgorithmResponseResourceRequirements", string(data)}, " ")
}
