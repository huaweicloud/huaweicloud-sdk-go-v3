package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmResponse 算法管理算法配置。
type AlgorithmResponse struct {
	Metadata *AlgorithmResponseMetadata `json:"metadata,omitempty"`

	JobConfig *AlgorithmResponseJobConfig `json:"job_config,omitempty"`

	// 算法资源约束，可不设置。设置后，在算法使用于训练作业时，控制台会过滤可用的公共资源池。
	ResourceRequirements *[]AlgorithmResponseResourceRequirements `json:"resource_requirements,omitempty"`

	AdvancedConfig *AlgorithmResponseAdvancedConfig `json:"advanced_config,omitempty"`
}

func (o AlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmResponse struct{}"
	}

	return strings.Join([]string{"AlgorithmResponse", string(data)}, " ")
}
