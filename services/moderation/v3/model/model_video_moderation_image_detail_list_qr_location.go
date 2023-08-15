package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VideoModerationImageDetailListQrLocation 二维码位置信息，该数组有四个值，分别代表左上角的坐标和右下角的坐标。例如[207,522,340,567] 207代表的是左上角的坐标 522代表左上角的坐标 340代表的是右下角的坐标 567代表的是右下角的坐标。
type VideoModerationImageDetailListQrLocation struct {

	// 检测出的二维码左上角横坐标。
	TopLeftX *int32 `json:"top_left_x,omitempty"`

	// 检测出的二维码左上角纵坐标。
	TopLeftY *int32 `json:"top_left_y,omitempty"`

	// 检测出的二维码右下角横坐标。
	BottomRightX *int32 `json:"bottom_right_x,omitempty"`

	// 检测出的二维码右下角纵坐标。
	BottomRightY *int32 `json:"bottom_right_y,omitempty"`
}

func (o VideoModerationImageDetailListQrLocation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VideoModerationImageDetailListQrLocation struct{}"
	}

	return strings.Join([]string{"VideoModerationImageDetailListQrLocation", string(data)}, " ")
}
