package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 目标检测框位置信息
type ImageTaggingBoundingBox struct {

	// 检测框区域宽度
	Width *float64 `json:"width,omitempty"`

	// 检测框区域高度
	Height *float64 `json:"height,omitempty"`

	// 检测框左上角到垂直轴距离
	TopLeftX *float64 `json:"top_left_x,omitempty"`

	// 检测框左上角到水平轴距离
	TopLeftY *float64 `json:"top_left_y,omitempty"`
}

func (o ImageTaggingBoundingBox) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageTaggingBoundingBox struct{}"
	}

	return strings.Join([]string{"ImageTaggingBoundingBox", string(data)}, " ")
}
