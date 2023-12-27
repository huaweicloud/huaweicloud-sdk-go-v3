package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayerPositionConfig 图层位置配置。
type LayerPositionConfig struct {

	// 图层左上角像素点的X轴位置值（画布左上角坐标是0x0）。 > * 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。
	Dx int32 `json:"dx"`

	// 图层图片左上角像素点的Y轴位置值（画布左上角坐标是0x0）。 > 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。
	Dy int32 `json:"dy"`

	// 图片/视频/人物图的层顺序。 > * 图层顺序从1开始的整数，底层图层顺序是1，往上依次增加。
	LayerIndex int32 `json:"layer_index"`
}

func (o LayerPositionConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerPositionConfig struct{}"
	}

	return strings.Join([]string{"LayerPositionConfig", string(data)}, " ")
}
