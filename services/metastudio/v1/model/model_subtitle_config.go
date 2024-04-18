package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubtitleConfig 字幕配置。
type SubtitleConfig struct {

	// 字幕框左下角像素点坐标。  > *横屏（16:9）视频像素为1920x1080；竖屏（9:16）视频像素为1080x1920。
	Dx *int32 `json:"dx,omitempty"`

	// 字幕框左下角像素点坐标。  > *横屏（16:9）视频像素为1920x1080；竖屏（9:16）视频像素为1080x1920。
	Dy *int32 `json:"dy,omitempty"`

	// 字体。当前支持的字体： * HarmonyOS_Sans_SC_Black：鸿蒙粗体 * HarmonyOS_Sans_SC_Regular：鸿蒙常规 * HarmonyOS_Sans_SC_Thin：鸿蒙细体
	FontName *string `json:"font_name,omitempty"`

	// 字体大小。  取值范围：[4, 120]
	FontSize *int32 `json:"font_size,omitempty"`

	// 字幕框高度 > * 参数h用于方便前端计算字幕框左上角坐标，后台不使用该参数
	H *int32 `json:"h,omitempty"`

	// 字幕框宽度 > * 字幕框宽度固定为屏幕宽度的80% > * 参数w用于方便前端计算字幕框左上角坐标，后台不使用该参数
	W *int32 `json:"w,omitempty"`
}

func (o SubtitleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtitleConfig struct{}"
	}

	return strings.Join([]string{"SubtitleConfig", string(data)}, " ")
}
