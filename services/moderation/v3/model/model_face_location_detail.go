package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FaceLocationDetail struct {

	// 检测出人脸的左上角横坐标。
	TopLeftX *int32 `json:"top_left_x,omitempty"`

	// 检测出人脸的左上角纵坐标。
	TopLeftY *int32 `json:"top_left_y,omitempty"`

	// 检测出人脸的右下角横坐标。
	BottomRightX *int32 `json:"bottom_right_x,omitempty"`

	// 检测出人脸的右下角纵坐标。
	BottomRightY *int32 `json:"bottom_right_y,omitempty"`
}

func (o FaceLocationDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FaceLocationDetail struct{}"
	}

	return strings.Join([]string{"FaceLocationDetail", string(data)}, " ")
}
