package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BoundingBox 结合口袋，包含口袋中心位置和尺寸大小
type BoundingBox struct {

	// 口袋中心坐标; x, y, z轴的坐标
	Center *[]float32 `json:"center,omitempty"`

	// 口袋尺寸大小;  x, y, z轴的大小
	Size *[]float32 `json:"size,omitempty"`
}

func (o BoundingBox) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BoundingBox struct{}"
	}

	return strings.Join([]string{"BoundingBox", string(data)}, " ")
}
