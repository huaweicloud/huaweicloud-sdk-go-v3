package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FepParamDto 自由能微扰任务的参数，预平衡时长、平衡时长和预平衡步数、平衡步数二者都存在时以预平衡时长、平衡时长为准。
type FepParamDto struct {

	// 预平衡时长，单位ps，范围为0-200，不包含0，默认为100ps
	PreEquilibriumTime *float32 `json:"pre_equilibrium_time,omitempty"`

	// 平衡时长，单位ns，范围为0-10，不包含0，默认为1ns
	EquilibriumTime *float32 `json:"equilibrium_time,omitempty"`

	// 预平衡步数，默认为50000
	NumPreEquilibriumSteps *int32 `json:"num_pre_equilibrium_steps,omitempty"`

	// 平衡步数，默认为500000
	NumEquilibriumSteps *int32 `json:"num_equilibrium_steps,omitempty"`

	// 时间步长，单位ps，取值范围：大于0，小于等于0.005
	TimestepSize *float32 `json:"timestep_size,omitempty"`

	// lambda个数
	NumLambda *int32 `json:"num_lambda,omitempty"`
}

func (o FepParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FepParamDto struct{}"
	}

	return strings.Join([]string{"FepParamDto", string(data)}, " ")
}
