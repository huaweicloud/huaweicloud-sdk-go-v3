package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PointCleanDto 点位清洗配置结构体
type PointCleanDto struct {

	// 静默时间窗口，在该时间窗口内，没有触发上报条件，点位将不会上
	SilentWindow int32 `json:"silent_window"`

	// 偏差，在该偏差范围内表示是正常波动，点位将不进行上报
	Deviation float64 `json:"deviation"`
}

func (o PointCleanDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PointCleanDto struct{}"
	}

	return strings.Join([]string{"PointCleanDto", string(data)}, " ")
}
