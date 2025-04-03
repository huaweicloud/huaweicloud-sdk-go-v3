package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayerRotationConfig 图层旋转配置。
type LayerRotationConfig struct {

	// **参数解释**： 旋转角度。 **取值范围**： 角度范围0-360度。 **默认取值**： 0度。 **约束限制**： 以素材中心点旋转。
	Angle *int32 `json:"angle,omitempty"`
}

func (o LayerRotationConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerRotationConfig struct{}"
	}

	return strings.Join([]string{"LayerRotationConfig", string(data)}, " ")
}
