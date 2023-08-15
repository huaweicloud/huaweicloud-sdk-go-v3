package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PointScalingDto 点位缩放配置结构体
type PointScalingDto struct {

	// 缩放的倍率
	Ratio float64 `json:"ratio"`

	// 基准值
	Base float64 `json:"base"`

	// 缩放后结果的精度，精确到小数点后几位,-1表示全部保留，0表示只保留整数位
	Accuracy *int32 `json:"accuracy,omitempty"`
}

func (o PointScalingDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PointScalingDto struct{}"
	}

	return strings.Join([]string{"PointScalingDto", string(data)}, " ")
}
