package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImageSpriteTaskOutPut struct {
	ObsInfo *ObsInfo `json:"obs_info,omitempty"`

	// 雪碧图的高度
	Height *int32 `json:"height,omitempty"`

	// 雪碧图的宽度
	Width *int32 `json:"width,omitempty"`

	// 雪碧图对象列表
	SpriteImageList *[]string `json:"sprite_image_list,omitempty"`
}

func (o ImageSpriteTaskOutPut) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageSpriteTaskOutPut struct{}"
	}

	return strings.Join([]string{"ImageSpriteTaskOutPut", string(data)}, " ")
}
