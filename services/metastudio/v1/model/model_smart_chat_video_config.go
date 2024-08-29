package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartChatVideoConfig struct {

	// 视频宽度。  单位：像素。  最小值320，最大值2560。 > * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280、3840x2160、2160x3840六种分辨率。4K分辨率视频需要分身数字人模型支持4K的情况下才能使用。 > * clip_mode=CROP，裁剪后视频，（dx,dy）为原点，保留视频像宽度为width。 > * 分身数字人直播目前只支持1080x1920。
	Width *int32 `json:"width,omitempty"`

	// 视频高度。  单位：像素。  最小值320，最大值2560。 > * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280、3840x2160、2160x3840六种分辨率分辨率。 > * clip_mode=CROP，裁剪后视频，（dx,dy）为原点，保留视频像高度为height。 > * 分身数字人直播目前只支持1080x1920。
	Height *int32 `json:"height,omitempty"`
}

func (o SmartChatVideoConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartChatVideoConfig struct{}"
	}

	return strings.Join([]string{"SmartChatVideoConfig", string(data)}, " ")
}
