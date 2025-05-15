package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FrozenEffect 冻结效果。
type FrozenEffect struct {
	FrozenEffect *FrozenEffectEnum `json:"frozen_effect,omitempty"`
}

func (o FrozenEffect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FrozenEffect struct{}"
	}

	return strings.Join([]string{"FrozenEffect", string(data)}, " ")
}
