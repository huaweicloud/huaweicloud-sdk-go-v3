package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomScalingConfig **参数描述**:  自定义扩容策略。  **约束限制**：  不涉及。
type CustomScalingConfig struct {
	EnlargeScene *ScalingScene `json:"enlarge_scene,omitempty"`
}

func (o CustomScalingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomScalingConfig struct{}"
	}

	return strings.Join([]string{"CustomScalingConfig", string(data)}, " ")
}
