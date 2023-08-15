package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FepParamDto struct {

	// 预平衡步数
	NumPreEquilibriumSteps *int32 `json:"num_pre_equilibrium_steps,omitempty"`

	// 平衡步数
	NumEquilibriumSteps *int32 `json:"num_equilibrium_steps,omitempty"`

	// 时间步长，取值范围:大于0，小于等于0.005
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
