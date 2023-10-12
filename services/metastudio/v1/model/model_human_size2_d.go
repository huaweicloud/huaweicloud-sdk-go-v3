package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HumanSize2D 分身数字人在背景图片中的大小。
type HumanSize2D struct {

	// 分身数字人宽度像素值。 > 横屏（16:9）背景图片像素为1920x1080；竖屏（9:16）背景图片像素为1080x1920。
	Width *int32 `json:"width,omitempty"`

	// 分身数字人高度像素值。 > 横屏（16:9）背景图片像素为1920x1080；竖屏（9:16）背景图片像素为1080x1920。
	Height *int32 `json:"height,omitempty"`
}

func (o HumanSize2D) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanSize2D struct{}"
	}

	return strings.Join([]string{"HumanSize2D", string(data)}, " ")
}
