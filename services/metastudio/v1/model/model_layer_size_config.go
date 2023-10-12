package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LayerSizeConfig 图层大小配置。
type LayerSizeConfig struct {

	// 图层图片宽度像素值（相对背景图片）。 > 横屏（16:9）背景图片像素为1920x1080；竖屏（9:16）背景图片像素为1080x1920。
	Width *int32 `json:"width,omitempty"`

	// 图层图片高度像素值（相对背景图片）。 > 横屏（16:9）背景图片像素为1920x1080；竖屏（9:16）背景图片像素为1080x1920。
	Height *int32 `json:"height,omitempty"`
}

func (o LayerSizeConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LayerSizeConfig struct{}"
	}

	return strings.Join([]string{"LayerSizeConfig", string(data)}, " ")
}
