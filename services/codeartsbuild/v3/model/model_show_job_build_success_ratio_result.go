package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobBuildSuccessRatioResult 构建成功率
type ShowJobBuildSuccessRatioResult struct {

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 分支
	Branch *string `json:"branch,omitempty"`

	// 构建成功总数
	TotalSuccessCount *int32 `json:"total_success_count,omitempty"`

	// 构建总数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 总成功比率分数
	TotalSuccessRatioFraction *string `json:"total_success_ratio_fraction,omitempty"`

	// 每日构建成功率
	EveryDayReport *[]ShowJobBuildSuccessRatioResultEveryDayReport `json:"every_day_report,omitempty"`
}

func (o ShowJobBuildSuccessRatioResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildSuccessRatioResult struct{}"
	}

	return strings.Join([]string{"ShowJobBuildSuccessRatioResult", string(data)}, " ")
}
