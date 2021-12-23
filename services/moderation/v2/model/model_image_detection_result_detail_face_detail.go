package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 涉政场景中的人物面部信息。
type ImageDetectionResultDetailFaceDetail struct {
	// 人脸区域高度。

	H *int32 `json:"h,omitempty"`
	// 人脸区域宽度。

	W *int32 `json:"w,omitempty"`
	// 人脸区域左上角到y轴距离。

	X *int32 `json:"x,omitempty"`
	// 人脸区域左上角到x轴距离。

	Y *int32 `json:"y,omitempty"`
}

func (o ImageDetectionResultDetailFaceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageDetectionResultDetailFaceDetail struct{}"
	}

	return strings.Join([]string{"ImageDetectionResultDetailFaceDetail", string(data)}, " ")
}
