package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StatUsageInfo struct {

	// 统计时间点，UTC时间，格式为YYYY-MM-DDTHH:MM:SSZ。
	StatTime *string `json:"stat_time,omitempty"`

	// 保底带宽。  说明： 该字段为预留值，当前始终为空；当场景为95增强时才返回数值。
	GuaranteedBandWidth *string `json:"guaranteed_band_width,omitempty"`

	// 用量。
	Usage *string `json:"usage,omitempty"`

	// 单位，您可以调用查询度量单位列表接口获取。带宽和用量使用相同的计量单位。
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o StatUsageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatUsageInfo struct{}"
	}

	return strings.Join([]string{"StatUsageInfo", string(data)}, " ")
}
