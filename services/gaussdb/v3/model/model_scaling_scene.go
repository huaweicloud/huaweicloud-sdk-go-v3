package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingScene **参数描述**：  变配场景。  **约束限制**：  只支持扩容场景。
type ScalingScene struct {

	// **参数解释**:  变配条件。  **约束限制**：  不涉及。
	Conditions *[]ScalingCondition `json:"conditions,omitempty"`
}

func (o ScalingScene) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingScene struct{}"
	}

	return strings.Join([]string{"ScalingScene", string(data)}, " ")
}
