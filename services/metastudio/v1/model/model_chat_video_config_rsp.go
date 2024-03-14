package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChatVideoConfigRsp struct {

	// 视频宽度。  单位：像素。  最小值320，最大值2560。 > * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280四种分辨率。
	Width *int32 `json:"width,omitempty"`

	// 视频高度。  单位：像素。  最小值320，最大值2560。 > * clip_mode=RESIZE时，当前支持1920x1080、1080x1920、1280x720、720x1280四种分辨率。
	Height *int32 `json:"height,omitempty"`
}

func (o ChatVideoConfigRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatVideoConfigRsp struct{}"
	}

	return strings.Join([]string{"ChatVideoConfigRsp", string(data)}, " ")
}
