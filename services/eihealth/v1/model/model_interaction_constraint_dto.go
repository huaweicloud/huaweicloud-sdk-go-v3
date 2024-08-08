package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InteractionConstraintDto 相互作用力约束参数
type InteractionConstraintDto struct {

	// 相互作用力列表
	Interactions []Interaction `json:"interactions"`

	// 是否排除指定的约束作用力
	Exclusive bool `json:"exclusive"`

	Operator *OperatorType `json:"operator,omitempty"`
}

func (o InteractionConstraintDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InteractionConstraintDto struct{}"
	}

	return strings.Join([]string{"InteractionConstraintDto", string(data)}, " ")
}
