package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 一个Region内的资源需求描述
type RegionSpec struct {

	// region的id
	RegionId string `json:"region_id"`

	// Region内期望的总算力容量
	ExpectTargetCapacity int32 `json:"expect_target_capacity"`

	// Region内期望的稳定算力总量
	ExpectStableCapacity int32 `json:"expect_stable_capacity"`

	LaunchTemplateConfig *LaunchTemplateConfig `json:"launch_template_config"`
}

func (o RegionSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegionSpec struct{}"
	}

	return strings.Join([]string{"RegionSpec", string(data)}, " ")
}
