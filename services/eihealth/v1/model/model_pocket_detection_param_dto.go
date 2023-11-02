package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PocketDetectionParamDto 靶点口袋发现参数列表
type PocketDetectionParamDto struct {

	// 时间步长，单位ps
	TimestepSize *float64 `json:"timestep_size,omitempty"`

	// 最小化步数
	NumMinimizationSteps *int32 `json:"num_minimization_steps,omitempty"`

	// 预平衡时长，单位ps
	PreEquilibriumTime *int32 `json:"pre_equilibrium_time,omitempty"`

	// 口袋发现时长，单位ns
	PocketDetectionTime *int32 `json:"pocket_detection_time,omitempty"`

	// 表面原子离散点数量
	NumSurfacePoints *int32 `json:"num_surface_points,omitempty"`

	// 探针半径，单位A
	ProbeRadius *float64 `json:"probe_radius,omitempty"`
}

func (o PocketDetectionParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PocketDetectionParamDto struct{}"
	}

	return strings.Join([]string{"PocketDetectionParamDto", string(data)}, " ")
}
