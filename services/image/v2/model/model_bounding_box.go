package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BoundingBox 目标检测框位置信息
type BoundingBox struct {

	// 检测框区域宽度
	Width *int32 `json:"width,omitempty"`

	// 检测框区域高度
	Height *int32 `json:"height,omitempty"`

	// 检测框左上角到垂直轴距离
	TopLeftX *int32 `json:"top_left_x,omitempty"`

	// 检测框左上角到水平轴距离
	TopLeftY *int32 `json:"top_left_y,omitempty"`
}

func (o BoundingBox) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BoundingBox struct{}"
	}

	return strings.Join([]string{"BoundingBox", string(data)}, " ")
}
