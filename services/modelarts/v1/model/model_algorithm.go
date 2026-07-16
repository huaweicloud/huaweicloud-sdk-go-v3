package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Algorithm 算法管理算法配置。
type Algorithm struct {
	Metadata *AlgorithmMetadata `json:"metadata,omitempty"`

	JobConfig *AlgorithmJobConfig `json:"job_config,omitempty"`

	// 算法资源约束。可不设置。设置后，在算法使用于训练作业时，控制台会过滤可用的公共资源池。
	ResourceRequirements *[]ResourceRequirement `json:"resource_requirements,omitempty"`

	AdvancedConfig *AlgorithmAdvancedConfig `json:"advanced_config,omitempty"`
}

func (o Algorithm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Algorithm struct{}"
	}

	return strings.Join([]string{"Algorithm", string(data)}, " ")
}
