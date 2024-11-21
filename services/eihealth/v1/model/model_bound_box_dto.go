package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BoundBoxDto 结合口袋，包含口袋中心位置和尺寸大小。
type BoundBoxDto struct {

	// 口袋中心坐标，x、y、z轴的坐标。
	Center []float64 `json:"center"`

	// 口袋尺寸大小，x、y、z轴的大小。
	Size []float32 `json:"size"`

	// 填充。
	Padding *float32 `json:"padding,omitempty"`
}

func (o BoundBoxDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BoundBoxDto struct{}"
	}

	return strings.Join([]string{"BoundBoxDto", string(data)}, " ")
}
