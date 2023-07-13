package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageFrame
type ImageFrame struct {

	// 创建时间
	CreateTime string `json:"create_time"`

	// 更新时间
	UpdateTime string `json:"update_time"`

	// 0: 16/9 1: 4/3 2: 3/4
	FrameResolutionType int32 `json:"frame_resolution_type"`

	//
	Id string `json:"id"`

	ImageFramePosition *LeftRightPosition `json:"image_frame_position"`

	ImagePosition *LeftRightPosition `json:"image_position"`

	// 横版：0 竖版：1
	ResolutionType int32 `json:"resolution_type"`

	//
	Url string `json:"url"`
}

func (o ImageFrame) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageFrame struct{}"
	}

	return strings.Join([]string{"ImageFrame", string(data)}, " ")
}
