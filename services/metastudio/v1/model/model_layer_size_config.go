package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayerSizeConfig 图层大小配置。
type LayerSizeConfig struct {

	// **参数解释**： 图层图片左上角像素点的Y轴位置值图层图片宽度像素值（相对画布大小）。 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。  **约束限制**： 该值为相对于画布的像素值，仅表示布局位置关系，与输出画面分辨率无关。
	Width *int32 `json:"width,omitempty"`

	// **参数解释**： 图层图片高度像素值（相对画布大小）。 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。  **约束限制**： 该值为相对于画布的像素值，仅表示布局位置关系，与输出画面分辨率无关。|
	Height *int32 `json:"height,omitempty"`
}

func (o LayerSizeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerSizeConfig struct{}"
	}

	return strings.Join([]string{"LayerSizeConfig", string(data)}, " ")
}
