package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegionSpec 一个区域内的资源需求描述
type RegionSpec struct {

	// 区域ID
	RegionId string `json:"region_id"`

	// 区域内期望的总算力容量
	ExpectTargetCapacity int32 `json:"expect_target_capacity"`

	// 区域内期望的稳定算力容量
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
