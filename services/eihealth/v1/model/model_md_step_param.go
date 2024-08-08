package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MdStepParam MD计算步骤配置
type MdStepParam struct {

	// 能量最小化的步骤
	EnergyMinimizationSteps *int32 `json:"energy_minimization_steps,omitempty"`

	// 等温等体步骤模拟的时长，单位ps
	Nvt *float32 `json:"nvt,omitempty"`

	// 等压等温步骤模拟的时长，单位ps
	Npt *float32 `json:"npt,omitempty"`

	// 平衡步骤模拟的时长，单位ns
	SimulationTime *float32 `json:"simulation_time,omitempty"`
}

func (o MdStepParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MdStepParam struct{}"
	}

	return strings.Join([]string{"MdStepParam", string(data)}, " ")
}
