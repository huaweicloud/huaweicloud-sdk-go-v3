package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackgroundImageConfig 背景图片大小及位置配置。
type BackgroundImageConfig struct {

	// **参数解释**： 背景图片左上角像素点的X轴位置值（画布左上角坐标是0x0）。 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。  **约束限制**： 需要保证背景图片要铺满画布。即dx <= 0，并且横屏时dx + width >=1920，竖屏时dx + width >=1080。
	Dx int32 `json:"dx"`

	// **参数解释**： 背景图片左上角像素点的Y轴位置值（画布左上角坐标是0x0）。 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。  **约束限制**：  需要保证背景图片要铺满画布。即dy <= 0，并且横屏时dy + height >=1080，竖屏时dy + height >=1920。
	Dy int32 `json:"dy"`

	// **参数解释**： 背景图片宽度像素值（相对画布大小）。 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。  **约束限制**： 需要保证背景图片要铺满画布。即width > 1080，并且横屏时dx + width >=1920，竖屏时dx + width >=1080。
	Width int32 `json:"width"`

	// **参数解释**： 背景图片高度像素值（相对画布大小）。 横屏（16:9）画布像素为1920x1080；竖屏（9:16）画布像素为1080x1920。  **约束限制**： 需要保证背景图片要铺满画布。height> 1080，并且横屏时dy + height >=1080，竖屏时dy + height >=1920。
	Height int32 `json:"height"`
}

func (o BackgroundImageConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackgroundImageConfig struct{}"
	}

	return strings.Join([]string{"BackgroundImageConfig", string(data)}, " ")
}
