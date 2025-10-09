package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerlessScalingPolicy **参数描述**：  查询Serverless自定义扩容策略响应体。  **约束限制**：  不涉及。
type ServerlessScalingPolicy struct {

	// **参数描述**：  自定义扩容步长。  **约束限制**：  不涉及。  **取值范围**：  2-算力上限的一半。  **默认取值**：  不涉及。
	EnlargeStepSize *string `json:"enlarge_step_size,omitempty"`

	CustomScalingConfig *CustomScalingConfig `json:"custom_scaling_config,omitempty"`
}

func (o ServerlessScalingPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerlessScalingPolicy struct{}"
	}

	return strings.Join([]string{"ServerlessScalingPolicy", string(data)}, " ")
}
