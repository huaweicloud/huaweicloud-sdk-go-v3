package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 结合口袋，包含口袋中心位置和尺寸大小
type BoundingBoxDto struct {

	// 口袋中心坐标; x, y, z轴的坐标
	Center []float32 `json:"center"`

	// 口袋尺寸大小;  x, y, z轴的大小
	Size []float32 `json:"size"`
}

func (o BoundingBoxDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BoundingBoxDto struct{}"
	}

	return strings.Join([]string{"BoundingBoxDto", string(data)}, " ")
}
