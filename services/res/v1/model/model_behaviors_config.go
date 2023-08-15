package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BehaviorsConfig
type BehaviorsConfig struct {

	// 正向行为。
	PositiveBehaviors *[]BehaviorWeights `json:"positive_behaviors,omitempty"`

	// 负向行为。
	NegativeBehaviors *[]BehaviorWeights `json:"negative_behaviors,omitempty"`
}

func (o BehaviorsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BehaviorsConfig struct{}"
	}

	return strings.Join([]string{"BehaviorsConfig", string(data)}, " ")
}
