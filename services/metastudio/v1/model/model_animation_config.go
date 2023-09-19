package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AnimationConfig 动作参数配置。
type AnimationConfig struct {

	// 动作资产ID。
	Animation *string `json:"animation,omitempty"`
}

func (o AnimationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AnimationConfig struct{}"
	}

	return strings.Join([]string{"AnimationConfig", string(data)}, " ")
}
