package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceUsageResponse Response Object
type ListResourceUsageResponse struct {

	// 有效天数，精度最高返回小数点后20位。  说明： 计算方式为上报的点数/288所得出的值。其中288为一天的点数，5分钟为一个点数单位。计算95费用时，因95费用是按月定价，若实际不足月，则是使用官网价*折扣*actual_days/当月天数，来计算费用明细。
	ActualDays *string `json:"actual_days,omitempty"`

	// 计费带宽的按月汇总。 说明： 每月2日20点后可查询上月数据；若查询当月数据，则为空。
	BandWidth *string `json:"band_width,omitempty"`

	// 月保底带宽的按月汇总。 说明： 每月2日20点后可查询上月数据；若查询当月数据，则为空。该字段为预留值，当前始终为空；当场景为95增强时才返回数值。
	MonthlyGuaranteedBandWidth *string `json:"monthly_guaranteed_band_width,omitempty"`

	// 月峰值带宽。 说明： 每月2日20点后可查询上月数据；若查询当月数据，则为空。该字段为预留值，当前始终为空；当场景为95增强时才返回数值。
	MonthlyPeakBandWidth *string `json:"monthly_peak_band_width,omitempty"`

	// 带宽单位，您可以调用查询度量单位列表接口获取。若所有带宽为空，则该字段为空。
	BandWidthMeasureId *int32 `json:"band_width_measure_id,omitempty"`

	// 总条数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 使用量明细（5分钟统计值）。具体请参见表3。
	UsageInfoList  *[]StatUsageInfo `json:"usage_info_list,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListResourceUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceUsageResponse struct{}"
	}

	return strings.Join([]string{"ListResourceUsageResponse", string(data)}, " ")
}
