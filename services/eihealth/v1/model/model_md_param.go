package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MdParam 靶点优化MD参数配置
type MdParam struct {

	// MD模拟的时间步长，单位fs，取值范围：大于0，小于等于5
	TimestepSize *float32 `json:"timestep_size,omitempty"`

	// MD模拟的温度，单位K
	Temperature *float32 `json:"temperature,omitempty"`

	StepParams *MdStepParam `json:"step_params,omitempty"`
}

func (o MdParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MdParam struct{}"
	}

	return strings.Join([]string{"MdParam", string(data)}, " ")
}
