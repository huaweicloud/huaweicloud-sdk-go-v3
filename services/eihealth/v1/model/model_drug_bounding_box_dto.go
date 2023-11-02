package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DrugBoundingBoxDto 结合口袋，包含口袋中心位置和尺寸大小
type DrugBoundingBoxDto struct {

	// 口袋中心坐标；x, y, z轴的坐标
	Center []float32 `json:"center"`

	// 口袋尺寸大小；x, y, z轴的大小
	Size []float32 `json:"size"`
}

func (o DrugBoundingBoxDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DrugBoundingBoxDto struct{}"
	}

	return strings.Join([]string{"DrugBoundingBoxDto", string(data)}, " ")
}
