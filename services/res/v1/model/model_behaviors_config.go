package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type BehaviorsConfig struct {

	// 正向行为。
	PositiveBehaviors *[]BehaviorWeights `json:"positive_behaviors,omitempty" xml:"positive_behaviors"`

	// 负向行为。
	NegativeBehaviors *[]BehaviorWeights `json:"negative_behaviors,omitempty" xml:"negative_behaviors"`
}

func (o BehaviorsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BehaviorsConfig struct{}"
	}

	return strings.Join([]string{"BehaviorsConfig", string(data)}, " ")
}
