package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlgorithmResponse Response Object
type CreateAlgorithmResponse struct {
	Metadata *AlgorithmResponseMetadata `json:"metadata,omitempty"`

	JobConfig *AlgorithmResponseJobConfig `json:"job_config,omitempty"`

	// 算法资源约束，可不设置。设置后，在算法使用于训练作业时，控制台会过滤可用的公共资源池。
	ResourceRequirements *[]AlgorithmResponseResourceRequirements `json:"resource_requirements,omitempty"`

	AdvancedConfig *AlgorithmResponseAdvancedConfig `json:"advanced_config,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o CreateAlgorithmResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlgorithmResponse struct{}"
	}

	return strings.Join([]string{"CreateAlgorithmResponse", string(data)}, " ")
}
