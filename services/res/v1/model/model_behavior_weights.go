package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BehaviorWeights struct {

	// 行为类型。
	BehaviorType *string `json:"behavior_type,omitempty"`

	// 行为权重。
	Weight *int32 `json:"weight,omitempty"`

	// 其他用途。
	OtherUses *[]string `json:"other_uses,omitempty"`
}

func (o BehaviorWeights) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BehaviorWeights struct{}"
	}

	return strings.Join([]string{"BehaviorWeights", string(data)}, " ")
}
