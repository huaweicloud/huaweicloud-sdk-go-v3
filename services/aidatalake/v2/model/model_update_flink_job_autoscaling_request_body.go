package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFlinkJobAutoscalingRequestBody **参数解释**：热更新自动调优配置参数请求体。 **约束限制**：不涉及。
type UpdateFlinkJobAutoscalingRequestBody struct {
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscaling_policy"`
}

func (o UpdateFlinkJobAutoscalingRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFlinkJobAutoscalingRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateFlinkJobAutoscalingRequestBody", string(data)}, " ")
}
